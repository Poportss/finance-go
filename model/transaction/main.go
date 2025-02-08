package transaction

import "time"

type Transaction struct {
	Title     string    `json:"title"`
	Value     float32   `json:"value"`
	Type      int       `json:"type"`
	CreatedAt time.Time `json:"created_at"`
}

type Transactions []Transaction
