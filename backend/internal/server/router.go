package server

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
	"github.com/nishojib/ffxivdailies/internal/api"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// RegisterRoutes returns a new http.Handler that routes requests to the correct handler.
func (s *Server) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	if s.cfg.Limiter.Enabled {
		router.Use(httprate.LimitByIP(s.cfg.Limiter.RPS, 1*time.Minute))
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
		v1Router.Get("/health", s.HealthHandler)
		v1Router.Post("/auth/login", s.LoginHandler)
		v1Router.Post("/auth/refresh", s.RefreshTokenHandler)
		v1Router.Post("/auth/revoke", s.RevokeTokenHandler)

		v1Router.Group(func(authRouter chi.Router) {
			authRouter.Use(s.withAuthentication)

			authRouter.Get("/tasks/shared", s.SharedTasksHandler)
			authRouter.Put("/tasks/shared/{taskID}", s.ToggleTaskHandler)
		})

	})

	return router
}
