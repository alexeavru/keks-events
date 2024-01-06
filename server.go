package main

import (
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alexeavru/keks-events/auth"
	"github.com/alexeavru/keks-events/database"
	"github.com/alexeavru/keks-events/graph"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
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

	router := chi.NewRouter()
	// Add Auth
	router.Use(auth.Middleware())
	// Enable CORS Support
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{http.MethodGet, http.MethodPost},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		EventsDB: eventsDb,
	}}))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/login", auth.CreateTokenEndpoint)

	log.Printf("Starting KEKS-events server ...")
	log.Printf("connect to http://0.0.0.0:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
