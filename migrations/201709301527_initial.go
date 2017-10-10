package main

import (
	"fmt"
	"github.com/go-pg/migrations"
	"github.com/odiak/MoneyForest/store"
)

func init() {
	migrations.Register(func(db migrations.DB) error {
		fmt.Println("creating initial tables")
		_, err := db.Exec(`
			CREATE EXTENSION "pgcrypto";

			CREATE TABLE users (
				id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
				email text NOT NULL UNIQUE,
				name text NOT NULL,
				encrypted_password text NOT NULL
			);

			CREATE TABLE account_types (
				id text PRIMARY KEY
			);

			CREATE TABLE accounts (
				id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
				owner_id uuid NOT NULL REFERENCES users (id),
				name text NOT NULL,
				description text NOT NULL,
				account_type text NOT NULL REFERENCES account_types (id)
					ON UPDATE CASCADE,
				balance integer NOT NULL,
				has_balance boolean
			);

			CREATE TABLE categories (
				id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
				owner_id uuid NOT NULL REFERENCES users (id),
				name text NOT NULL,
				parent_category_id uuid REFERENCES categories (id)
			);

			CREATE TABLE transaction_types (
				id text PRIMARY KEY
			);

			CREATE TABLE transactions (
				id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
				account_id uuid NOT NULL REFERENCES accounts (id),
				amount integer,
				title text NOT NULL,
				original_title text NOT NULL,
				description text NOT NULL,
				category_id uuid REFERENCES categories (id),
				date date NOT NULL,
				transaction_type text NOT NULL REFERENCES transaction_types (id),
				created_at timestamptz,
				updated_at timestamptz
			);
		`)
		if err != nil {
			return err
		}
		err = db.Insert(
			store.AccountType{"bank"},
			store.AccountType{"credit-card"},
			store.AccountType{"wallet"},
		)
		if err != nil {
			return err
		}
		err = db.Insert(
			store.TransactionType{"expense"},
			store.TransactionType{"income"},
			store.TransactionType{"transfer"},
			store.TransactionType{"balance-adjustment"},
		)
		if err != nil {
			return err
		}
		return nil
	}, func(db migrations.DB) error {
		fmt.Println("dropping initial tables")
		_, err := db.Exec(`
			DROP TABLE IF EXISTS users, account_types, accounts, categories, transactions CASCADE;
			DROP EXTENSION IF EXISTS "pgcrypto";
		`)
		return err
	})
}
