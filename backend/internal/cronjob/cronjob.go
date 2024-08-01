package cronjob

import (
	"log/slog"
	"time"

	"github.com/r3labs/sse/v2"
	"github.com/robfig/cron/v3"
)

func New(publisher messagePublisher) *cron.Cron {
	c := cron.New(cron.WithSeconds(), cron.WithLocation(time.UTC))

	slog.Info("Starting cron job to signal for page refresh every tuesday at 8am UTC (3am EST)")
	id, err := c.AddFunc("0 0 8 * * 2", func() {
		slog.Info("Tasks reset successfully for tuesday")
		publisher.Publish("messages", &sse.Event{Data: []byte("Weekly tasks reset")})
	})
	if err != nil {
		slog.Error("failed to add cron job", "cronjobID", id, "error", err)
	}

	slog.Info("Starting cron job to signal for page refresh every day at 3pm UTC (11am EST)")
	id, err = c.AddFunc("0 0 15 * * *", func() {
		slog.Info("Tasks reset successfully for every day")
		publisher.Publish("messages", &sse.Event{Data: []byte("Daily tasks reset")})
	})
	if err != nil {
		slog.Error("failed to add cron job", "cronjobID", id, "error", err)
	}

	return c
}

type messagePublisher interface {
	Publish(id string, event *sse.Event)
}
