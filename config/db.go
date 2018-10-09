package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)


// Database config object
func DB () *sql.DB {
	db, err := sql.Open("sqlite3", "./app.db")

	if err != nil {
		panic(err)
	}

	return db
}