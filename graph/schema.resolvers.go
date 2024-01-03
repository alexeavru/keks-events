package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/alexeavru/keks-events/graph/model"
)

// User is the resolver for the user field.
func (r *eventResolver) User(ctx context.Context, obj *model.Event) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	randNumber, _ := rand.Int(rand.Reader, big.NewInt(100))
	event := &model.Event{
		DateStart: input.DateStart,
		EventName: input.EventName,
		Text:      input.Text,
		ID:        fmt.Sprintf("T%d", randNumber),
		UserID:    input.UserID,
	}
	r.events = append(r.events, event)
	return event, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	return r.events, nil
}

// Event returns EventResolver implementation.
func (r *Resolver) Event() EventResolver { return &eventResolver{r} }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type eventResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
