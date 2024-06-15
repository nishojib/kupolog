package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/nishojib/ffxivdailies/docs"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/server"
	"github.com/nishojib/ffxivdailies/internal/vcs"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type config struct {
	port       string
	db         db
	env        api.Environment
	limiter    api.Limiter
	version    string
	authSecret string
}

type db struct {
	dsn          string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  time.Duration
}

// TODO: change this to bearer when PR https://github.com/swaggo/swag/pull/1821 is merged

//	@title						Swagger Kupolog API
//	@version					1.0
//	@description				This is an API for the Kupolog app.
//	@termsOfService				https://api.kupolog.com/terms
//
//	@contact.name				nishojib
//	@contact.url				https://api.kupolog.com/support
//	@contact.email				nishojib@kupolog.com
//
//	@license.name				MIT
//	@license.url				https://opensource.org/license/mit
//
//	@BasePath					/v1
//
//	@securitydefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization
//	@description				"Type 'Bearer TOKEN' to correctly set the API Key"

// @externalDocs.description	OpenAPI
// @externalDocs.url			https://swagger.io/resources/open-api/
func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	envStr := flag.String("env", "development", "Environment")
	flag.Parse()

	env, err := api.NewEnvironment(*envStr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	// if err := godotenv.Load(); err != nil && env.String() == api.EnvDevelopment {
	// 	slog.Error(err.Error())
	// 	os.Exit(1)
	// }

	docs.SwaggerInfo.Host = os.Getenv("API_URL")

	cfg := config{
		env:        env,
		port:       os.Getenv("PORT"),
		db:         db{dsn: os.Getenv("DB_DSN")},
		version:    vcs.Version(),
		authSecret: os.Getenv("AUTH_SECRET"),
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
		cfg.authSecret,
		server.WithPort(cfg.port),
	)

	if err := s.Serve(cfg.env); err != nil {
		return nil
	}

	return nil
}

func (cfg *config) initDB() (*bun.DB, error) {
	dsn := strings.Split(os.Getenv("DB_DSN"), "?")[0]
	slog.Info("connecting to db...", "dsn", dsn)

	sqldb, err := sql.Open("libsql", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	sqldb.SetMaxOpenConns(cfg.db.maxOpenConns)
	sqldb.SetMaxIdleConns(cfg.db.maxIdleConns)
	sqldb.SetConnMaxIdleTime(cfg.db.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqldb.PingContext(ctx); err != nil {
		err := sqldb.Close()
		return nil, err
	}

	slog.Info("connected to db...")

	db := bun.NewDB(sqldb, sqlitedialect.New())
	if cfg.env == api.EnvDevelopment {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return db, nil
}
