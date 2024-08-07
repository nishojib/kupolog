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
	"github.com/nishojib/ffxivdailies/internal/task"
	"github.com/nishojib/ffxivdailies/internal/user"
	"github.com/r3labs/sse/v2"
)

// Server represents an HTTP server.
type Server struct {
	srv       *http.Server
	db        Repository
	userModel UserModel
	authModel AuthModel
	sse       *sse.Server
	provider  Provider
	clients   map[chan string]bool
	wg        sync.WaitGroup

	cfg Config
}

type Config struct {
	Limiter    api.Limiter
	Env        api.Environment
	Version    string
	AuthSecret string
}

// New creates a new server with the provided option.Options.
func New(
	db Repository,
	userModel UserModel,
	authModel AuthModel,
	provider Provider,
	sse *sse.Server,
	cfg Config,
) *Server {
	s := &Server{
		db:        db,
		userModel: userModel,
		authModel: authModel,
		provider:  provider,
		sse:       sse,
		cfg:       cfg,
		clients:   make(map[chan string]bool),
	}

	s.srv = &http.Server{
		Handler:      s.RegisterRoutes(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(slog.Default().Handler(), slog.LevelError),
	}

	return s
}

// Serve starts the server and blocks until the server is stopped.
func (s *Server) ListenAndServe(port int) error {

	shutdownError := make(chan error, 1)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		sig := <-quit

		slog.Info("shutting down server", "signal", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err := s.srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		slog.Info("completing background tasks", "addr", s.srv.Addr)

		s.wg.Wait()
		shutdownError <- nil
	}()

	s.srv.Addr = fmt.Sprintf(":%d", port)
	err := s.srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	slog.Info("stopped server", "addr", s.srv.Addr)
	return nil
}

// Repository is an interface that represents the db operations for the server.

//go:generate mockery --with-expecter --name Repository
type Repository interface {
	GetUserByProviderID(ctx context.Context, providerAccountID string) (user.User, error)
	InsertAndLinkAccount(ctx context.Context, u *user.User, account *user.Account) error
	GetUserByUserID(ctx context.Context, userID string) (user.User, error)

	IsTokenRevoked(ctx context.Context, token string) (bool, error)
	RevokeToken(ctx context.Context, token string) error

	AddUserTask(ctx context.Context, t *task.Task) error
	UpdateUserTask(ctx context.Context, t *task.Task) error
	GetUserTask(ctx context.Context, userID string, taskID string) (task.Task, error)
	GetTasksForUser(ctx context.Context, userID string) ([]task.Task, error)
}

// Provider is an interface that represents the auth provider.

//go:generate mockery --with-expecter --name Provider
type Provider interface {
	Validate(provider string, token string) (string, bool, error)
}

// UserModel is an interface that represents the user model.

//go:generate mockery --with-expecter --name UserModel
type UserModel interface {
	GetOrCreate(
		ctx context.Context,
		email user.Email,
		provider user.Provider,
		accountID user.ID,
	) (user.User, error)
}

// AuthModel is an interface that represents the auth model.

//go:generate mockery --with-expecter --name AuthModel
type AuthModel interface {
	CreateTokens(
		ctx context.Context,
		userID string,
		authSecret string,
	) (string, string, error)
	RefreshToken(token, secret string) (string, error)
	GetBearerToken(headers http.Header) (string, error)
	IsTokenRevoked(ctx context.Context, token string) (bool, error)
}
