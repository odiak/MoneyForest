package store

import (
	"time"
)

type Transaction struct {
	ID              string `sql:"type:uuid"`
	AccountID       string `sql:"type:uuid"`
	Account         *Account
	Amount          int32
	Title           string
	OriginalTitle   string
	Description     string
	CategoryID      string `sql:"type:uuid"`
	Category        *Category
	Date            time.Time
	TransactionType string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
