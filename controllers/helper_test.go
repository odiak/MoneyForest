package controllers

import (
	"testing"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/odiak/MoneyForest/store"
)

func runInTx(t *testing.T, db *pg.DB, f func(orm.DB)) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			t.Error(err)
		}
	}()

	f(tx)
	err = tx.Rollback()
	if err != nil {
		t.Error(err)
	}
}

func createUser(t *testing.T, db orm.DB, email string) *store.User {
	user := &store.User{
		Name:  "test user",
		Email: email,
	}
	user.SetPassword("password")
	err := db.Insert(user)
	if err != nil {
		t.Error(err)
	}
	return user
}

func createAccount(t *testing.T, db orm.DB, userID, name, accountType string) *store.Account {
	account := &store.Account{
		OwnerID:     userID,
		Name:        name,
		Description: "",
		AccountType: accountType,
		Balance:     0,
		HasBalance:  true,
	}
	err := db.Insert(account)
	if err != nil {
		t.Error(err)
	}
	return account
}

func createCategory(t *testing.T, db orm.DB, userID string, parentID *string, name string) *store.Category {
	category := &store.Category{
		OwnerID:          userID,
		ParentCategoryID: parentID,
		Name:             name,
	}
	err := db.Insert(category)
	if err != nil {
		t.Error(err)
	}
	return category
}
