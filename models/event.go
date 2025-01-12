package models

import (
	"time"

	"github.com/chriskoorzen/go-rest-demo/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"       binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"    binding:"required"`
	DateTime    time.Time `json:"datetime"    binding:"required"`
	UserID      int       `json:"user_id"`
}

func (event *Event) Save() error {
	// Save event to database
	query := `
	INSERT INTO events (title, description, location, datetime, userID)
	VALUES (?, ?, ?, ?, ?)`

	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmnt.Close()

	result, err := stmnt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	event.ID = id

	return err
}

func GetAllEvents() []Event {
	// Get all events from database
	return []Event{}
}
