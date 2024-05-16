package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/google"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/server"
	"github.com/nishojib/ffxivdailies/internal/vcs"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type config struct {
	port    string
	db      db
	google  provider
	env     api.Environment
	limiter api.Limiter
	version string
}

type provider struct {
	key      string
	secret   string
	callback string
}

type db struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	if err := godotenv.Load(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	env, err := api.NewEnvironment(os.Getenv("ENVIRONMENT"))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	cfg := config{
		env:  env,
		port: os.Getenv("PORT"),
		db:   db{dsn: os.Getenv("DB_DSN")},
		google: provider{
			key:      os.Getenv("GOOGLE_KEY"),
			secret:   os.Getenv("GOOGLE_SECRET"),
			callback: os.Getenv("GOOGLE_CALLBACK_URL"),
		},
		version: vcs.Version(),
	}

	flag.IntVar(&cfg.db.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&cfg.db.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(
		&cfg.db.maxIdleTime,
		"db-max-idle-time",
		15*time.Minute,
		"PostgreSQL max connection idle time",
	)
	flag.IntVar(&cfg.limiter.RPS, "limiter-rps", 100, "Rate limiter requests per second")
	flag.BoolVar(&cfg.limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()

	goth.UseProviders(google.New(cfg.google.key, cfg.google.secret, cfg.google.callback))

	if err := cfg.run(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func (cfg *config) run() error {
	db, err := cfg.initDB()
	if err != nil {
		return err
	}

	s := server.New(
		db,
		cfg.limiter,
		cfg.env,
		cfg.version,
		server.WithPort(cfg.port),
	)

	if err := s.Serve(cfg.env); err != nil {
		return nil
	}

	return nil
}

func (cfg *config) initDB() (*sql.DB, error) {
	slog.Info("connecting to db...")

	db, err := sql.Open("libsql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConns)
	db.SetMaxIdleConns(cfg.db.maxIdleConns)
	db.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}

	slog.Info("connected to db...")

	return db, nil
}
