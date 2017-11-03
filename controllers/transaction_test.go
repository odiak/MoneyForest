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
	"github.com/odiak/MoneyForest/constants"
	"github.com/odiak/MoneyForest/util"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransaction(t *testing.T) {
	var (
		service = goa.New("MoneyForest-test")
		db      = pg.Connect(config.PgOptions)
	)

	runInTx(t, db, func(db orm.DB) {
		var (
			ctrl      = NewTransactionController(service, db)
			ctx       = context.Background()
			user      = createUser(t, db, "user1@example.com")
			account   = createAccount(t, db, user.ID, "test account", "bank")
			category1 = createCategory(t, db, user.ID, nil, "category1")
			category2 = createCategory(t, db, user.ID, util.StringPtr(category1.ID), "category2")
		)

		ctx = context.WithValue(ctx, constants.CurrentUserKey, user)

		t.Run("successful", func(t *testing.T) {
			payload := &app.TransactionPayload{
				Title:           "title 1",
				OriginalTitle:   "title 1",
				Amount:          300,
				Description:     "",
				AccountID:       uuid.FromStringOrNil(account.ID),
				CategoryID:      util.UUIDPtr(uuid.FromStringOrNil(category2.ID)),
				TransactionType: "expense",
				Date:            "2017-10-30",
			}

			_, media := test.CreateTransactionOK(t, ctx, service, ctrl, payload)

			if media.Amount != 300 {
				t.Error("wrong amount")
			}

			if media.Category == nil || media.Category.ID.String() != category2.ID {
				t.Error("wrong category")
			}
			if media.Category.ParentCategory != nil && media.Category.ParentCategory.ID.String() != category1.ID {
				t.Error("wrong parent category")
			}
		})
	})
}
