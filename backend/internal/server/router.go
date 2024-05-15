package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/server/handlers"
)

// NewRoutes returns a new http.Handler that routes requests to the correct handler.
func NewRoutes(
	db *sql.DB,
	limiter api.Limiter,
	env api.Environment,
	version string,
) http.Handler {
	router := chi.NewRouter()

	if limiter.Enabled {
		router.Use(httprate.LimitByIP(limiter.RPS, 1*time.Minute))
	}

	router.Use(middleware.Recoverer)

	router.Mount("/debug", middleware.Profiler())

	router.NotFound(api.NotFoundResponse)
	router.MethodNotAllowed(api.MethodNotAllowedResponse)

	router.Get("/health", handlers.Health(env, version))

	return router
}
