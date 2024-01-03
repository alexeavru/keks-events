package database

import (
	"database/sql"
)

type Event struct {
	db        *sql.DB
	ID        string
	EventName string
	Text      string
}

func NewEvent(db *sql.DB) *Event {
	return &Event{db: db}
}

func (c *Event) FindAll() ([]Event, error) {
	rows, err := c.db.Query("SELECT id, event_name, text FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []Event{}
	for rows.Next() {
		var id, event_name, text string
		if err := rows.Scan(&id, &event_name, &text); err != nil {
			return nil, err
		}
		events = append(events, Event{ID: id, EventName: event_name, Text: text})
	}
	return events, nil
}
