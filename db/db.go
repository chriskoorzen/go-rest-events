package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB // export pointer to database instance

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
	createUsersTable := `
    CREATE TABLE IF NOT EXISTS users (
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      email TEXT NOT NULL UNIQUE,
      password TEXT NOT NULL
    );`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println("Error creating 'Users' table")
		panic(err)
	}

	createEventsTable := `
    CREATE TABLE IF NOT EXISTS events (
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      title TEXT NOT NULL,
      description TEXT NOT NULL,
      location TEXT NOT NULL,
      datetime TEXT NOT NULL,
      creatorID INTEGER,
      FOREIGN KEY (creatorID) REFERENCES users(id)
    );`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		fmt.Println("Error creating 'Events' table")
		panic(err)
	}

	createRegistrationsTable := `
    CREATE TABLE IF NOT EXISTS registrations (
      id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
      eventID INTEGER NOT NULL,
      userID INTEGER NOT NULL,
      FOREIGN KEY (eventID) REFERENCES events(id),
      FOREIGN KEY (userID) REFERENCES users(id),
	  UNIQUE(eventID, userID)
    );`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		fmt.Println("Error creating 'Registrations' table")
		panic(err)
	}
}
