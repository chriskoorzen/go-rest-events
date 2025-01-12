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
}
