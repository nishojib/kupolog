package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/options"
	"github.com/uptrace/bun"
)

// Server represents an HTTP server.
type Server struct {
	*http.Server
	wg sync.WaitGroup
}

// New creates a new server with the provided option.Options.
func New(
	db *bun.DB,
	limiter api.Limiter,
	env api.Environment,
	version string,
	authSecret string,
	opts ...options.Option[Server],
) *Server {
	s := &Server{
		Server: &http.Server{
			Addr:         ":8080",
			Handler:      NewRoutes(db, limiter, env, version, authSecret),
			IdleTimeout:  1 * time.Minute,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			ErrorLog:     slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
		},
	}

	return s.Append(opts...)
}

// Serve starts the server and blocks until the server is stopped.
func (s *Server) Serve(env api.Environment) error {
	shutdownError := make(chan error, 1)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		slog.Info("shutting down server", "signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := s.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		slog.Info("completing background tasks", "addr", s.Addr)

		s.wg.Wait()
		shutdownError <- nil
	}()

	slog.Info("starting server", "addr", s.Addr, "env", env.String())

	err := s.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	slog.Info("stopped server", "addr", s.Addr)

	return nil

}

// Append applies the provided option.Options to the server.
func (s *Server) Append(opts ...options.Option[Server]) *Server {
	for _, opt := range opts {
		opt.Apply(s)
	}

	return s
}

// WithPort sets the address for the server.
func WithPort(port string) options.Option[Server] {
	return options.OptionFunc[Server](func(s *Server) {
		s.Server.Addr = fmt.Sprintf(":%s", port)
	})
}
