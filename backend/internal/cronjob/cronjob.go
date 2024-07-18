package cronjob

import (
	"context"
	"log/slog"
	"time"

	"github.com/r3labs/sse/v2"
	"github.com/robfig/cron/v3"
)

func New(taskResetter taskResetter, publisher messagePublisher) *cron.Cron {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	c := cron.New(cron.WithSeconds(), cron.WithLocation(time.UTC))

	slog.Info("Starting cron job to reset tasks every tuesday at 8am UTC (3am EST)")
	c.AddFunc("0 0 8 * * 2", func() {
		err := resetTasksEveryTuesday(ctx, taskResetter)
		if err != nil {
			slog.Error("failed to reset tasks", "error", err)
			return
		}
		slog.Info("Tasks reset successfully for tuesday")
		publisher.Publish("messages", &sse.Event{Data: []byte("Weekly tasks reset")})

	})

	slog.Info("Starting cron job to reset tasks every day at 3pm UTC (11am EST)")
	c.AddFunc("0 0 15 * * *", func() {
		err := resetTasksEveryday(ctx, taskResetter)
		if err != nil {
			slog.Error("failed to reset tasks", "error", err)
			return
		}
		slog.Info("Tasks reset successfully for every day")
		publisher.Publish("messages", &sse.Event{Data: []byte("Daily tasks reset")})
	})

	return c
}

func resetTasksEveryTuesday(ctx context.Context, db taskResetter) error {
	err := db.UpdateTaskForKind(ctx, "weekly")
	return err
}

func resetTasksEveryday(ctx context.Context, db taskResetter) error {
	err := db.UpdateTaskForKind(ctx, "daily")
	return err
}

type taskResetter interface {
	UpdateTaskForKind(ctx context.Context, kind string) error
}

type messagePublisher interface {
	Publish(id string, event *sse.Event)
}
