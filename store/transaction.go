package store

import (
	"time"
)

type Transaction struct {
	ID              string `sql:"type:uuid"`
	AccountID       string `sql:"type:uuid,notnull"`
	Account         *Account
	Amount          int32  `sql:",notnull"`
	Title           string `sql:",notnull"`
	OriginalTitle   string `sql:",notnull"`
	Description     string `sql:",notnull"`
	CategoryID      string `sql:"type:uuid,notnull"`
	Category        *Category
	Date            time.Time `sql:",notnull"`
	TransactionType string    `sql:",notnull"`
	CreatedAt       time.Time `sql:",notnull"`
	UpdatedAt       time.Time `sql:",notnull"`
}
