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