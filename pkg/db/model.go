package db

type Coin struct {
	ID         string `json:"ID"`
	UnvID      string `json:"UnvID"`
	Name       string `json:"name"`
	Symbol     string `json:"symbol"`
	Price      string `json:"price"`
	LastSynced string `json:"lastSynced"`
	Tracked    bool   `json:"tracked"`
}
