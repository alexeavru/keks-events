package handlers

import (
	"net/http"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/alexeavru/keks-events/internal/auth"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

// Router register necessary routes and returns an instance of a router.
func Router(srv *handler.Server) chi.Router {

	isReady := &atomic.Value{}
	// Set readyz probe is negative by default
	isReady.Store(false)

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

	// Set readyz probe is positive
	isReady.Store(true)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)
	router.HandleFunc("/login", auth.CreateTokenEndpoint)
	router.HandleFunc("/healthz", Healthz)
	router.HandleFunc("/readyz", Readyz(isReady))

	return router
}

// healthz is a liveness probe.
func Healthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// readyz is a readiness probe.
func Readyz(isReady *atomic.Value) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		if isReady == nil || !isReady.Load().(bool) {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}
