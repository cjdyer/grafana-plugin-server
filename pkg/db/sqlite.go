package db

import (
	"embed"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

var DB *sqlx.DB

func Init() error {
	var err error
	DB, err = sqlx.Open("sqlite3", "./plugins.db")
	if err != nil {
		return err
	}

	if err := runMigrations(); err != nil {
		return err
	}

	return DB.Ping()
}

func runMigrations() error {
	schema, err := migrationFiles.ReadFile("migrations/schema.sql")
	if err != nil {
		return err
	}
	_, err = DB.Exec(string(schema))
	return err
}
