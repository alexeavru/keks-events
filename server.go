package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alexeavru/keks-events/graph"
)

const defaultPort = "8080"

func EventsCatalog(t *testing.T) {
	client, err := Open("sqlite3", "file:event-catalog.db?cache=shared&_fk=1")

	assert.NoErrorf(t, err, "failed opening connection to sqlite")
	defer client.Close()

	ctx := context.Background()

	// Run the automatic migration tool to create all schema resources.
	err = client.Schema.Create(ctx)
	assert.NoErrorf(t, err, "failed creating schema resources")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
