package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type EventDB struct {
	db          *sql.DB
	ID          string
	EventName   string
	Description string
	DateStart   string
	DateEnd     string
}

func NewEvent(db *sql.DB) *EventDB {
	return &EventDB{db: db}
}

func (c *EventDB) FindAll() ([]EventDB, error) {
	rows, err := c.db.Query("SELECT id, event_name, description, date_start, date_end FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []EventDB{}
	for rows.Next() {
		var id, event_name, description, date_start, date_end string
		if err := rows.Scan(&id, &event_name, &description, &date_start, &date_end); err != nil {
			return nil, err
		}
		events = append(events, EventDB{ID: id, EventName: event_name, Description: description, DateStart: date_start, DateEnd: date_end})
	}
	return events, nil
}

func (c *EventDB) Create(event_name string, description string, date_start string, date_end string) (EventDB, error) {
	id := uuid.New().String()
	//=====================================
	// Пока захардкордим
	//=====================================
	user_id, priority, group := 0, 0, 0
	//=====================================
	_, err := c.db.Exec("INSERT INTO events ('id', 'event_name', 'description', 'date_start', 'date_end', 'user_id', 'priority', 'group') VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		id, event_name, description, date_start, date_end, user_id, priority, group)
	if err != nil {
		return EventDB{}, err
	}
	return EventDB{ID: id, EventName: event_name, Description: description, DateStart: date_start, DateEnd: date_end}, nil
}

func (c *EventDB) Delete(id string) (bool, error) {
	result, err := c.db.Exec("DELETE FROM events WHERE id=?", id)
	status, _ := result.RowsAffected()
	if status == 0 {
		return false, err
	}
	return true, nil
}
