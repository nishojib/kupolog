package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/repository"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	envStr := flag.String("env", "development", "Environment")
	kind := flag.String("kind", "daily", "Kind of tasks to update")
	flag.Parse()

	env, err := api.NewEnvironment(*envStr)
	if err != nil {
		slog.Error("failed to get environment", "error", err)
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil && env.String() == api.EnvDevelopment {
		slog.Error("failed to load env", "error", err)
		os.Exit(1)
	}

	bunDB, err := initDB(env)
	if err != nil {
		slog.Error("failed to init db", "error", err)
		os.Exit(1)
	}

	repo := repository.New(bunDB)

	if *kind != "daily" && *kind != "weekly" {
		slog.Error("invalid kind: kind must be either weekly or daily", "kind", *kind)
		os.Exit(1)
	}

	err = repo.UpdateTaskForKind(context.Background(), *kind)
	if err != nil {
		slog.Error("failed to update tasks", "error", err)
		os.Exit(1)
	}

	slog.Info(fmt.Sprintf("%s tasks updated successfully", *kind))
}

func initDB(env api.Environment) (*bun.DB, error) {
	dsn := os.Getenv("DB_DSN")

	dsnUrl := strings.Split(os.Getenv("DB_DSN"), "?")[0]
	slog.Info("connecting to db...", "dsn", dsnUrl)

	sqldb, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, err
	}

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
