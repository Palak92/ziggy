package db

import (
	"database/sql"
	"fmt"
	"log"
	"sort"

	"github.com/go-sql-driver/mysql"
	"github.com/palak92/ziggy/pkg/coinmarketcap"
)

type CoinsDB struct {
	DB *sql.DB
}

func New(user, password, address string) (*CoinsDB, error) {
	//Capture connection properties.
	cfg := mysql.Config{
		User:                 user,
		Passwd:               password,
		Net:                  "tcp",
		Addr:                 address,
		DBName:               "coins",
		AllowNativePasswords: true,
	}
	// Get a database handle.

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("err while opening mysql %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("err while pinging mysql %v", err)
	}

	cDB := &CoinsDB{
		DB: db,
	}
	err = cDB.initData()
	if err != nil {
		return nil, fmt.Errorf("error while initializing data:%w", err)
	}
	return cDB, nil
}

func (c *CoinsDB) Close() {
	c.DB.Close()
}
func (c *CoinsDB) deleteAllCoins() error {
	// Prepare a DELETE statement to remove all rows from the table
	stmt, err := c.DB.Prepare("DELETE FROM coins")
	if err != nil {
		return fmt.Errorf("error while preparing statement: %w", err)
	}
	defer stmt.Close()

	// Execute the statement to remove all rows from the table
	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("error while executing statement: %w", err)
	}

	log.Default().Println("All rows removed from table.")
	return nil
}

func (c *CoinsDB) initData() error {
	// clean database
	err := c.deleteAllCoins()
	if err != nil {
		return fmt.Errorf("error while deleting all rows of coins table:%w", err)
	}

	list, err := coinmarketcap.ListCoins(5)
	if err != nil {
		return fmt.Errorf("error while getting list from coinmarketcap:%w", err)
	}
	for _, crypto := range list.Data {
		coin := &Coin{
			UnvID:      fmt.Sprint(crypto.ID),
			Name:       crypto.Name,
			Symbol:     crypto.Symbol,
			Tracked:    true,
			Price:      fmt.Sprint(crypto.Quote["USD"].Price),
			LastSynced: crypto.LastUpdated,
		}
		err = c.InsertCoin(coin)
		if err != nil {
			return fmt.Errorf("error while adding coin in database:%w", err)
		}
	}
	log.Default().Println("5 coins added to database")
	return nil
}

func (c *CoinsDB) InsertCoin(coin *Coin) error {
	// Prepare the INSERT statement
	stmt, err := c.DB.Prepare("INSERT INTO coins (coin_id, name, symbol, tracked,price, last_synced) VALUES (?, ?, ?,?,?,?)")
	if err != nil {
		return fmt.Errorf("error while preparing statement: %w", err)
	}
	defer stmt.Close()

	// Execute the INSERT statement
	res, err := stmt.Exec(coin.UnvID, coin.Name, coin.Symbol, coin.Tracked, coin.Price, coin.LastSynced)
	if err != nil {
		return fmt.Errorf("error while executing statement: %w", err)
	}

	// Get the last inserted ID
	id, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("error while getting last id: %w", err)
	}
	log.Default().Printf("Inserted user with ID %d\n", id)
	return nil
}

func (c *CoinsDB) AllCoins() ([]*Coin, error) {
	// An coins slice to hold data from returned rows.
	var coins []*Coin

	rows, err := c.DB.Query("SELECT * FROM coins")
	if err != nil {
		return nil, fmt.Errorf("err while querying all coins: %w", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var coin Coin
		if err := rows.Scan(&coin.ID, &coin.UnvID, &coin.Name, &coin.Symbol, &coin.Tracked, &coin.Price, &coin.LastSynced); err != nil {
			return nil, fmt.Errorf("err while scanning coin from query: %w", err)
		}
		coins = append(coins, &coin)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("err from rows received from query: %w", err)
	}
	return coins, nil
}

// CoinsByName queries for coins that have the specified  name.
func (c *CoinsDB) CoinsByName(name string) (*Coin, error) {
	var coins []*Coin

	rows, err := c.DB.Query("SELECT * FROM coins WHERE name = ?", name)
	if err != nil {
		return nil, fmt.Errorf("while running query to get coins by name %q: %v", name, err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var coin Coin
		if err := rows.Scan(&coin.ID, &coin.UnvID, &coin.Name, &coin.Symbol, &coin.Tracked, &coin.Price, &coin.LastSynced); err != nil {
			return nil, fmt.Errorf("err while scanning coin from query: %w", err)
		}
		coins = append(coins, &coin)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("err from rows received from query: %w", err)
	}

	coins = sortCoins(coins, "dcs")

	return coins[0], nil
}

func (c *CoinsDB) TrackCoin(coinID string, track bool) error {
	// Prepare the update statement
	stmt, err := c.DB.Prepare("UPDATE coins SET tracked = ? WHERE coin_id = ?")
	if err != nil {
		fmt.Errorf("error while preparing statement: %w", err)
	}
	defer stmt.Close()

	// Execute the update statement
	res, err := stmt.Exec(track, coinID)
	if err != nil {
		fmt.Errorf("error while executing update: %w", err)
	}

	// Check how many rows were affected by the update
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		fmt.Errorf("error while getting rows affected: %w", err)
	}

	log.Default().Printf("Updated %d rows\n", rowsAffected)
	return nil
}
func sortCoins(coins []*Coin, order string) []*Coin {
	if order == "asc" {
		sort.Slice(coins, func(i, j int) bool {
			return coins[i].LastSynced > coins[j].LastSynced
		})
	}
	sort.Slice(coins, func(i, j int) bool {
		return coins[i].LastSynced < coins[j].LastSynced
	})
	return coins
}

func (c *CoinsDB) GetTrackedCoinIDs() ([]string, error) {
	var coinIDs []string

	rows, err := c.DB.Query("SELECT DISTINCT coin_id FROM coins where tracked = true")
	if err != nil {
		return nil, fmt.Errorf("error while querying unique coin Ids: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var coinID string
		if err := rows.Scan(&coinID); err != nil {
			return nil, fmt.Errorf("error while scanning coin Id: %w", err)
		}
		coinIDs = append(coinIDs, coinID)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error while getting coin Ids: %w", err)
	}
	return coinIDs, nil
}
