package db

import (
	"database/sql"
	"fmt"
	"log"

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
	defer db.Close()

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

func (c *CoinsDB) initData() error {
	list, err := coinmarketcap.ListCoins(5)
	if err != nil {
		return fmt.Errorf("error while getting list from coinmarketcap:%w", err)
	}
	for _, crypto := range list.CryptoMarket {
		coin := &Coin{
			UnvID:      fmt.Sprint(crypto.ID),
			Name:       crypto.Name,
			Symbol:     crypto.Symbol,
			Tracked:    false,
			Price:      fmt.Sprint(crypto.Quote["USD"].Price),
			LastSynced: crypto.LastUpdated,
		}
		err = c.addCoins(coin)
		if err != nil {
			return fmt.Errorf("error while adding coin in database:%w", err)
		}
	}
	log.Default().Println("5 coins added to database")
	return nil
}

func (c *CoinsDB) addCoins(coin *Coin) error {
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

func (c *CoinsDB) AllCoins() ([]Coin, error) {
	// An coins slice to hold data from returned rows.
	var coins []Coin

	rows, err := c.DB.Query("SELECT * FROM coins = ?")
	if err != nil {
		return nil, fmt.Errorf("err while querying all coins: %w", err)
	}
	defer rows.Close()
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var coin Coin
		if err := rows.Scan(&coin.ID, &coin.UnvID, &coin.Name, &coin.Symbol, &coin.Price, &coin.LastSynced, &coin.Tracked); err != nil {
			return nil, fmt.Errorf("err while scanning coin from query: %w", err)
		}
		coins = append(coins, coin)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("err from rows received from query: %w", err)
	}
	return coins, nil
}
