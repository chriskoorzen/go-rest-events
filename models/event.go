package models

import (
	"time"

	"github.com/chriskoorzen/go-rest-events/db"
)

type Event struct {
	ID          int64     `json:"-"`
	Title       string    `json:"title"       binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"    binding:"required"`
	DateTime    time.Time `json:"datetime"    binding:"required"`
	CreatorID   int64     `json:"-"`
}

func (event *Event) Save() error {
	// Save event to database
	query := `
	INSERT INTO events (title, description, location, datetime, creatorID)
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
		event.CreatorID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()

	event.ID = id

	return err
}

func (event *Event) Update() error {
	// Update event in database
	query := `
	UPDATE events
	SET title = ?, description = ?, location = ?, datetime = ?
	WHERE id = ?`

	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmnt.Close()
	_, err = stmnt.Exec(
		event.Title,
		event.Description,
		event.Location,
		event.DateTime.Format(time.RFC3339Nano),
		event.ID,
	)

	return err
}

func (event Event) Delete() error {
	// Delete event from database
	query := "DELETE FROM events WHERE id = ?"
	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()

	_, err = stmnt.Exec(event.ID)

	return err
}

func GetAllEvents() ([]Event, error) {
	// Get all events from database
	query := `
	SELECT title, description, location, datetime
	FROM events`
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
			&event.Title,
			&event.Description,
			&event.Location,
			&datetimeStr,
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
	query := `
	SELECT title, description, location, datetime
	FROM events
	WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event
	var datetimeStr string
	err := row.Scan(
		&event.Title,
		&event.Description,
		&event.Location,
		&datetimeStr,
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

func (event *Event) Register(userID int64) error {
	query := `
	INSERT INTO registrations (eventID, userID)
	VALUES (?, ?)`

	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()

	_, err = stmnt.Exec(event.ID, userID)

	return err
}

func (event Event) CancelRegistration(userID int64) error {
	query := `
	DELETE FROM registrations
	WHERE eventID = ? AND userID = ?`

	stmnt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmnt.Close()

	_, err = stmnt.Exec(event.ID, userID)

	return err
}

func (event *Event) GetRegistrations() ([]UserExposedJSON, error) {
	query := `
	SELECT u.email
	FROM registrations r
	JOIN users u ON r.userID = u.id
	WHERE r.eventID = ?`

	rows, err := db.DB.Query(query, event.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []UserExposedJSON
	for rows.Next() {
		var user UserExposedJSON
		err := rows.Scan(&user.Email)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
