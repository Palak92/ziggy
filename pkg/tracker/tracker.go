// Package tracker tracks the price of coin in every 5 minutes.
package tracker

import (
	"fmt"
	"log"
	"time"

	"github.com/palak92/ziggy/pkg/coinmarketcap"
	"github.com/palak92/ziggy/pkg/db"
)

type Tracker struct {
	DB     *db.CoinsDB
	ticker *time.Ticker
	done   chan bool
}

func New(db *db.CoinsDB) *Tracker {
	return &Tracker{
		DB:     db,
		ticker: time.NewTicker(15 * time.Second),
		done:   make(chan bool),
	}
}
func (t Tracker) Start() {
	fmt.Println("Tracker started")
	go func() {
		for {
			select {
			case <-t.done:
				return
			case <-t.ticker.C:
				t.addCoinPrices()
			}
		}
	}()

}

func (t Tracker) Stop() {
	t.ticker.Stop()
	t.done <- true
	fmt.Println("Tracker stopped")
}

func (t Tracker) addCoinPrices() {
	coinIDs, err := t.DB.GetTrackedCoinIDs()
	if err != nil {
		log.Default().Printf("error while getting coin Ids: %v", err)
		return
	}

	if len(coinIDs) == 0 {
		log.Default().Printf("No coins are tracked!!")
		return
	}

	log.Default().Printf("getting latest quotes for coin ids %v", coinIDs)
	quotes, err := coinmarketcap.GetLatestQuotes(coinIDs)
	if err != nil {
		log.Fatalf("error while getting latest quotes of coin Ids: %v", err)
		return
	}

	for _, q := range quotes.Data {
		c := &db.Coin{
			UnvID:      fmt.Sprint(q.ID),
			Name:       q.Name,
			Symbol:     q.Symbol,
			Price:      fmt.Sprint(q.Quote["USD"].Price),
			LastSynced: q.LastUpdated,
			Tracked:    true,
		}
		err := t.DB.InsertCoin(c)
		if err != nil {
			log.Fatalf("error while inserting coin %v in coins table: %v", c, err)
			return
		}
	}

}
