//go:generate goagen bootstrap -d github.com/odiak/MoneyForest/design

package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/odiak/MoneyForest/app"
	"github.com/odiak/MoneyForest/config"
	"github.com/odiak/MoneyForest/constants"
	"github.com/odiak/MoneyForest/controllers"
	"github.com/odiak/MoneyForest/store"
)

func main() {
	// Create service
	service := goa.New("MoneyForest")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	db := pg.Connect(config.PgOptions)
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		service.LogInfo(fmt.Sprintf("SQL Query: %s, %s", time.Since(event.StartTime), query))
	})

	app.UseAPIKeyAuthMiddleware(service, NewAPIKeyMiddleware(db))
	app.MountUserController(service, controllers.NewUserController(service, db))
	app.MountAccountController(service, controllers.NewAccountController(service, db))
	app.MountCategoryController(service, controllers.NewCategoryController(service, db))

	// Start service
	if err := service.ListenAndServe(fmt.Sprintf(":%d", config.Port)); err != nil {
		service.LogError("startup", "err", err)
	}

}

func NewAPIKeyMiddleware(db orm.DB) goa.Middleware {
	scheme := app.NewAPIKeyAuthSecurity()

	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			key := req.Header.Get(scheme.Name)
			if len(key) == 0 {
				return goa.ErrUnauthorized("missing auth token")
			}
			ut := store.UserToken{}
			err := db.Model(&ut).
				Column("user_token.*", "User").
				Where("user_token.token = ?", key).
				Select()
			if err != nil {
				if err == pg.ErrNoRows {
					return goa.ErrUnauthorized("invalid auth token")
				}
				return goa.ErrInternal("unknown error")
			}
			goa.LogInfo(ctx, "valid auth token", "token", key)
			ctx = context.WithValue(ctx, constants.CurrentUserKey, ut.User)
			return h(ctx, rw, req)
		}
	}
}
