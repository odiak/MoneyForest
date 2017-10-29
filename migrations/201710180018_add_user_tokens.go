package main

import (
	"fmt"
	"github.com/go-pg/migrations"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating user_tokens")
		_, err := db.Exec(`
			CREATE TABLE user_tokens (
				token text PRIMARY KEY,
				user_id uuid NOT NULL REFERENCES users (id) ON DELETE CASCADE,
				created_at timestamptz NOT NULL
			);
		`)
		return err
	}, func(db migrations.DB) error {
		fmt.Println("dropping user_tokens")
		_, err := db.Exec(`
			DROP TABLE IF EXISTS user_tokens CASCADE;
		`)
		return err
	})
}
