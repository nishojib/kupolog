package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/server/handlers"
	httpSwagger "github.com/swaggo/http-swagger/v2"
	"github.com/uptrace/bun"
)

// NewRoutes returns a new http.Handler that routes requests to the correct handler.
func NewRoutes(
	db *bun.DB,
	limiter api.Limiter,
	env api.Environment,
	version string,
	authSecret string,
) http.Handler {
	router := chi.NewRouter()

	if limiter.Enabled {
		router.Use(httprate.LimitByIP(limiter.RPS, 1*time.Minute))
	}

	router.Use(middleware.Logger, middleware.Recoverer, cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.NotFound(api.NotFoundResponse)
	router.MethodNotAllowed(api.MethodNotAllowedResponse)
	router.Mount("/debug", middleware.Profiler())

	router.Route("/v1", func(v1Router chi.Router) {
		v1Router.Mount("/swagger", httpSwagger.WrapHandler)
		v1Router.Get("/health", handlers.Health(env, version))
		v1Router.Post("/users", handlers.AddUser(db))
		v1Router.Post("/auth/login", handlers.Login(db, authSecret))
		v1Router.Post("/auth/refresh", handlers.RefreshToken(db, authSecret))
		v1Router.Post("/auth/revoke", handlers.RevokeToken(db))
	})

	return router
}
