package main

import (
	"context"
	"database/sql"
	"flag"
	"log/slog"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/joho/godotenv"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/task"
	"github.com/nrednav/cuid2"
	_ "github.com/tursodatabase/libsql-client-go/libsql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dbfixture"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	envStr := flag.String("env", "development", "Environment")
	flag.Parse()

	env, err := api.NewEnvironment(*envStr)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if err = godotenv.Load(); err != nil && env.String() == api.EnvDevelopment {
		slog.Error(err.Error())
		os.Exit(1)
	}

	dsn := strings.Split(os.Getenv("DB_DSN"), "?")[0]
	slog.Info("connecting to db...", "dsn", dsn)

	sqldb, err := sql.Open("libsql", os.Getenv("DB_DSN"))
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err = sqldb.PingContext(ctx); err != nil {
		err = sqldb.Close()
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("connected to db...")

	db := bun.NewDB(sqldb, sqlitedialect.New())
	if env == api.EnvDevelopment {
		db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))
	}

	db.RegisterModel((*task.Task)(nil))

	funcMap := template.FuncMap{
		"cuid2": func() string {
			return cuid2.Generate()
		},
		"now": func() string {
			return time.Now().Format(time.RFC3339Nano)
		},
	}

	slog.Info("seeding tasks...")
	fixture := dbfixture.New(db, dbfixture.WithTemplateFuncs(funcMap))
	if err = fixture.Load(ctx, os.DirFS("testdata"), "tasks.yml"); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	slog.Info("seeding completed...")
}
