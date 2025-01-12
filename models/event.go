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
	result, err := stmnt.Exec(
		event.Title,
		event.Description,
		event.Location,
		event.DateTime.Format(time.RFC3339Nano),
		event.UserID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	event.ID = id

	return err
}

func GetAllEvents() ([]Event, error) {
	// Get all events from database
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		var datetimeStr string
		err := rows.Scan(
			&event.ID,
			&event.Title,
			&event.Description,
			&event.Location,
			&datetimeStr,
			&event.UserID,
		)
		if err != nil {
			return nil, err
		}
		event.DateTime, err = time.Parse(time.RFC3339Nano, datetimeStr)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	// Get single event by ID
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	var datetimeStr string
	err := row.Scan(
		&event.ID,
		&event.Title,
		&event.Description,
		&event.Location,
		&datetimeStr,
		&event.UserID,
	)
	if err != nil {
		return nil, err
	}
	event.DateTime, err = time.Parse(time.RFC3339Nano, datetimeStr)
	if err != nil {
		return nil, err
	}

	return &event, nil
}
