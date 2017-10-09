package store

import (
	"time"
)

type Transaction struct {
	ID              string
	AccountID       string
	Account         *Account
	Amount          int32
	Title           string
	OriginalTitle   string
	Description     string
	CategoryID      string
	Category        *Category
	Date            time.Time
	TransactionType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
