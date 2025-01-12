package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"       binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location"    binding:"required"`
	DateTime    time.Time `json:"datetime"    binding:"required"`
	UserID      int       `json:"user_id"`
}

func (e Event) Save() error {
	// Save event to database
	return nil
}

func GetAllEvents() []Event {
	// Get all events from database
	return []Event{}
}
