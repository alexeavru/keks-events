package main

import (
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/alexeavru/keks-events/graph"
	"github.com/alexeavru/keks-events/internal/database"
	"github.com/alexeavru/keks-events/internal/handlers"
	"github.com/alexeavru/keks-events/internal/version"
	"github.com/joho/godotenv"
)

const defaultPort = "8088"

// init is invoked before main()
func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	database.InitDB()
	eventsDb := database.NewEvent(database.Db)
	defer database.CloseDB()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		EventsDB: eventsDb,
	}}))

	router := handlers.Router(srv)

	log.Printf("Starting KEKS-events server. Release %s BuildTime: %s Commit: %s", version.Release, version.BuildTime, version.Commit)
	log.Printf("connect to http://0.0.0.0:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
