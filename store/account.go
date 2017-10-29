package store

import (
	"github.com/go-pg/pg/orm"
)

type Account struct {
	ID          string `sql:"type:uuid"`
	OwnerID     string `sql:"type:uuid,notnull"`
	Owner       User
	Name        string `sql:",notnull"`
	Description string `sql:",notnull"`
	AccountType string `sql:",notnull"`
	Balance     int32  `sql:",notnull"`
	HasBalance  bool   `sql:",notnull"`
}

func (a *Account) validate() error {
	if a.Name == "" {
		return ValidationError("name is required")
	}
	return nil
}

func (a *Account) BeforeInsert(db orm.DB) error {
	return a.validate()
}

func (a *Account) BeforeUpdate(db orm.DB) error {
	return a.validate()
}
