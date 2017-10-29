package controllers

import (
	"context"
	"testing"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goadesign/goa"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/app/test"
	"github.com/odiak/MoneyForest/config"
	"github.com/odiak/MoneyForest/store"
)

func TestRegister(t *testing.T) {
	var (
		service = goa.New("MoneyForest-test")
		db      = pg.Connect(config.PgOptions)
	)

	runInTx(db, func(db orm.DB) {
		var (
			ctrl = NewUserController(service, db)
			ctx  = context.Background()
		)

		t.Run("successful", func(t *testing.T) {
			payload := &app.UserPayload{
				Email:    "foo@example.com",
				Name:     "JJJ XXX",
				Password: "1234abcd",
			}
			_, userMedia := test.RegisterUserOK(t, ctx, service, ctrl, payload)

			if userMedia.Email != "foo@example.com" {
				t.Error("wrong email")
			}
			if userMedia.Name != "JJJ XXX" {
				t.Error("wrong name")
			}

			user := store.User{}
			err := db.Model(&user).Where("email = ?", userMedia.Email).Select()
			if err != nil {
				t.Error(err)
			}
			if !user.ValidPassword("1234abcd") {
				t.Error("wrong password")
			}
		})

		t.Run("duplicated email", func(t *testing.T) {
			payload := &app.UserPayload{
				Email:    "foo@example.com",
				Name:     "JJJ XXX",
				Password: "1234abcd",
			}
			test.RegisterUserBadRequest(t, ctx, service, ctrl, payload)
		})

		t.Run("short password", func(t *testing.T) {
			payload := &app.UserPayload{
				Email:    "bar@example.com",
				Name:     "JJJ XXX",
				Password: "1234",
			}
			test.RegisterUserBadRequest(t, ctx, service, ctrl, payload)
		})
	})
}
