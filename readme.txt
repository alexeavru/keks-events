## Init 
go mod init github.com/alexeavru/keks-events

## create a tools.go
//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
)

## To automatically add the dependency to your go.mod
go mod tidy

## Import gqlgen
go get -d github.com/99designs/gqlgen@v0.17.42

## Create the project skeleton
go run github.com/99designs/gqlgen init

## Define your schema
graph/schema.graphqls

## Implement the resolvers
graph/schema.resolvers.go

## First let’s enable autobind
uncommenting the autobind config line in gqlgen.yml

## lets put it in graph/resolver.go
type Resolver struct {
	todos []*model.Event
}

## resolver.go, between package and import
//go:generate go run github.com/99designs/gqlgen generate

## run generate 
go generate ./...

## Run server
go run server.go


## Создать DB
sqlite3 events.db
create table events(id text primary key, event_name text, text text, user_id int, priority int, group int, date_start text, date_end text);

INSERT INTO events VALUES('7915ee64-87cd-4aba-a27e-27fe3759b412', 'Новый год', 'Отпраздновать новый год', '1', '0','0','2023-01-01T00:00:00Z','2023-01-08T23:59:59Z');
select * from events;

mutation createEvent {
  createEvent(input: { dateStart: "2023-01-02T00:00:00Z", eventName: "First event", text: "Event text", userId: "3" }) {
    user {
      id
    }
    dateStart
    eventName
    text
  }
}

query findEvents {
  events {
    eventName
    dateStart
    text
    user {
      name
    }
  }
}