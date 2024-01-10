package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.42

import (
	"context"
	"log"

	"github.com/alexeavru/keks-events/graph/model"
	"github.com/alexeavru/keks-events/internal/auth"
)

// CreateEvent is the resolver for the createEvent field.
func (r *mutationResolver) CreateEvent(ctx context.Context, input model.NewEvent) (*model.Event, error) {
	// Get UserID from context
	rawCtx := auth.ForContext(ctx)

	event, err := r.EventsDB.Create(input, rawCtx.UserID)
	if err != nil {
		return nil, err
	}

	log.Printf("Create New Event ID: %s", event.ID)

	return &model.Event{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Start:       input.Start,
		End:         input.End,
	}, nil
}

// UpdateEvent is the resolver for the updateEvent field.
func (r *mutationResolver) UpdateEvent(ctx context.Context, input model.UpdateEvent) (*model.Event, error) {
	// Get UserID from context
	rawCtx := auth.ForContext(ctx)

	event, err := r.EventsDB.Update(input, rawCtx.UserID)
	if err != nil {
		return nil, err
	}

	log.Printf("Update Event ID: %s", event.ID)

	return &model.Event{
		ID:          input.ID,
		Title:       input.Title,
		Description: input.Description,
		Start:       input.Start,
		End:         input.End,
	}, nil
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, id string) (bool, error) {
	status, err := r.EventsDB.Delete(id)
	if err != nil {
		return status, err
	}
	log.Printf("Delete event ID: %s", id)
	return status, nil
}

// Events is the resolver for the events field.
func (r *queryResolver) Events(ctx context.Context) ([]*model.Event, error) {
	// Get UserID from context
	rawCtx := auth.ForContext(ctx)

	eventsDB, err := r.EventsDB.FindAll(rawCtx.UserID)
	if err != nil {
		return nil, err
	}

	var eventsModel []*model.Event
	for _, event := range eventsDB {
		eventsModel = append(eventsModel, &model.Event{
			ID:          event.ID,
			Title:       event.Title,
			Description: event.Description,
			Start:       event.Start,
			End:         event.End,
		})
	}
	log.Printf("Get All Events")
	return eventsModel, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
