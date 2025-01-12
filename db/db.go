package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./db.sqlite")

	if err != nil {
		fmt.Println("Error opening database")
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
    id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    datetime TEXT NOT NULL,
    userID INTEGER
);`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println("Error creating 'Events' table")
		panic(err)
	}
}
