package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() error {
	var err error
	DB, err = sql.Open("sqlite3", "./plugins.db")
	if err != nil {
		return err
	}

	if err := runMigrations(); err != nil {
		return err
	}

	return DB.Ping()
}

func runMigrations() error {
	return nil
}
