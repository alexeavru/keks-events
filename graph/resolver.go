package graph

import (
	"github.com/alexeavru/keks-events/database"
	"github.com/alexeavru/keks-events/graph/model"
)

//go:generate go run github.com/99designs/gqlgen generate

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	EventsDB *database.Event
	events   []*model.Event
}
