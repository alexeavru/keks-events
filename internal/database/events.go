package database

import (
	"database/sql"
	"log"

	"github.com/alexeavru/keks-events/graph/model"
	"github.com/google/uuid"
)

type EventDB struct {
	db          *sql.DB
	ID          string
	Title       string
	Description string
	Start       string
	End         string
}

var Db *sql.DB

func InitDB() {
	db, err := sql.Open("sqlite", "./events.db")
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func NewEvent(db *sql.DB) *EventDB {
	return &EventDB{db: db}
}

func (c *EventDB) FindAll(user_id string) ([]EventDB, error) {
	rows, err := c.db.Query("SELECT id, title, description, start, end FROM events WHERE user_id = $1 ", user_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []EventDB{}
	for rows.Next() {
		var id, title, description, start, end string
		if err := rows.Scan(&id, &title, &description, &start, &end); err != nil {
			return nil, err
		}
		events = append(events, EventDB{ID: id, Title: title, Description: description, Start: start, End: end})
	}
	return events, nil
}

func (c *EventDB) Create(input model.NewEvent, user_id string) (EventDB, error) {
	id := uuid.New().String()
	//=====================================
	// Пока захардкордим
	//=====================================
	priority, group := 0, 0
	//=====================================
	_, err := c.db.Exec("INSERT INTO events ('id', 'title', 'description', 'start', 'end', 'user_id', 'priority', 'group') VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		id, input.Title, input.Description, input.Start, input.End, user_id, priority, group)
	if err != nil {
		return EventDB{}, err
	}
	return EventDB{ID: id, Title: input.Title, Description: input.Description, Start: input.Start, End: input.End}, nil
}

func (c *EventDB) Update(input model.UpdateEvent, user_id string) (EventDB, error) {
	_, err := c.db.Exec("UPDATE events SET title = $2, description = $3, start = $4, end = $4 WHERE id = $1",
		input.ID, input.Title, input.Description, input.Start, input.End)
	if err != nil {
		return EventDB{}, err
	}
	return EventDB{ID: input.ID, Title: input.Title, Description: input.Description, Start: input.Start, End: input.End}, nil
}

func (c *EventDB) Delete(id string) (bool, error) {
	result, err := c.db.Exec("DELETE FROM events WHERE id=?", id)
	status, _ := result.RowsAffected()
	if status == 0 {
		return false, err
	}
	return true, nil
}
