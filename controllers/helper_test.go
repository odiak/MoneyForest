package controllers

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

func runInTx(db *pg.DB, f func(orm.DB)) {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			panic(err)
		}
	}()

	f(tx)
	err = tx.Rollback()
	if err != nil {
		panic(err)
	}
}
