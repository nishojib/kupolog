package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/nishojib/ffxivdailies/docs"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/cronjob"
	"github.com/nishojib/ffxivdailies/internal/provider"
	"github.com/nishojib/ffxivdailies/internal/repository"
	"github.com/nishojib/ffxivdailies/internal/server"
	"github.com/nishojib/ffxivdailies/internal/vcs"
	"github.com/r3labs/sse/v2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type dbConfig struct {
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

	if err := godotenv.Load(); err != nil && env.String() == api.EnvDevelopment {
		slog.Error(err.Error())
		os.Exit(1)
	}

	docs.SwaggerInfo.Host = os.Getenv("API_URL")

	var dbCfg dbConfig

	dbCfg.dsn = os.Getenv("DB_DSN")

	flag.IntVar(&dbCfg.maxOpenConns, "db-max-open-conns", 25, "PostgreSQL max open connections")
	flag.IntVar(&dbCfg.maxIdleConns, "db-max-idle-conns", 25, "PostgreSQL max idle connections")
	flag.DurationVar(
		&dbCfg.maxIdleTime,
		"db-max-idle-time",
		15*time.Minute,
		"PostgreSQL max connection idle time",
	)

	var limiter api.Limiter
	flag.IntVar(&limiter.RPS, "limiter-rps", 100, "Rate limiter requests per second")
	flag.BoolVar(&limiter.Enabled, "limiter-enabled", true, "Enable rate limiter")

	flag.Parse()

	db, err := initDB(dbCfg, env)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	repo := repository.New(db)

	sseServer := sse.New()
	defer sseServer.RemoveStream("messages")
	defer sseServer.Close()

	sseServer.CreateStream("messages")

	c := cronjob.New(repo, sseServer)
	c.Start()
	defer c.Stop()

	port, err := strconv.Atoi(os.Getenv("PORT"))

	s := server.New(
		repo,
		provider.New(),
		sseServer,
		server.Config{
			Limiter:    limiter,
			Env:        env,
			Version:    vcs.Version(),
			AuthSecret: os.Getenv("AUTH_SECRET"),
		},
	)

	slog.Info("starting server", "addr", port, "env", env.String())

	if err != nil {
		slog.Error("failed to parse port", "error", err)
		os.Exit(1)
	}

	err = s.ListenAndServe(port)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}

func initDB(dbCfg dbConfig, env api.Environment) (*bun.DB, error) {
	dsn := strings.Split(os.Getenv("DB_DSN"), "?")[0]
	slog.Info("connecting to db...", "dsn", dsn)

	sqldb, err := sql.Open("libsql", dbCfg.dsn)
	if err != nil {
		return nil, err
	}

	sqldb.SetMaxOpenConns(dbCfg.maxOpenConns)
	sqldb.SetMaxIdleConns(dbCfg.maxIdleConns)
	sqldb.SetConnMaxIdleTime(dbCfg.maxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := sqldb.PingContext(ctx); err != nil {
		err := sqldb.Close()
		return nil, err
	}

	slog.Info("connected to db...")

	bunDB := bun.NewDB(sqldb, sqlitedialect.New())
	if env == api.EnvDevelopment {
		bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	return bunDB, nil
}
