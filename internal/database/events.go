package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/alexeavru/keks-events/graph/model"
	"github.com/google/uuid"
	"github.com/jackc/pgx"
)

type EventDB struct {
	db          *pgx.Conn
	ID          string
	Title       string
	Description string
	Start       string
	End         string
}

var Db *pgx.Conn

func InitDB() {
	cfg := pgx.ConnConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Database: os.Getenv("DATABASE_DB"),
		User:     os.Getenv("DATABASE_USER"),
		Password: os.Getenv("DATABASE_PASS"),
		RuntimeParams: map[string]string{
			"standard_conforming_strings": "on",
		},
		PreferSimpleProtocol: true,
	}
	db, err := pgx.Connect(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	Db = db
}

func CloseDB() error {
	return Db.Close()
}

func NewEvent(db *pgx.Conn) *EventDB {
	return &EventDB{db: db}
}

func (c *EventDB) FindAll(user_id string) ([]EventDB, error) {
	rows, err := c.db.Query("SELECT id, title, description, start, \"end\" FROM events WHERE user_id = $1", user_id)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	events := []EventDB{}

	for rows.Next() {
		var id, title, description, start, end string
		err := rows.Scan(&id, &title, &description, &start, &end)
		if err != nil {
			if err != sql.ErrNoRows {
				log.Print(err)
			}
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
	_, err := c.db.Exec("INSERT INTO events (id, title, description, start, \"end\", user_id, priority, \"group\") VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
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
	result, err := c.db.Exec("DELETE FROM events WHERE id = $1", id)
	status := result.RowsAffected()
	if status == 0 {
		return false, err
	}
	return true, nil
}
