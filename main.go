//go:generate goagen bootstrap -d github.com/odiak/MoneyForest/design

package main

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/odiak/MoneyForest/app"
)

func main() {
	// Create service
	service := goa.New("MoneyForest")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	db := pg.Connect(&pg.Options{
		User:     "kaido",
		Addr:     "127.0.0.1:5432",
		Database: "money_forest",
	})
	db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
		query, err := event.FormattedQuery()
		if err != nil {
			panic(err)
		}
		service.LogInfo(fmt.Sprintf("SQL Query: %s, %s", time.Since(event.StartTime), query))
	})

	// Mount "user" controller
	c := NewUserController(service, db)
	app.MountUserController(service, c)

	// Start service
	if err := service.ListenAndServe(":8000"); err != nil {
		service.LogError("startup", "err", err)
	}

}
