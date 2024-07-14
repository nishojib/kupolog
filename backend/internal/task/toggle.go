package task

import (
	"context"
	"errors"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"golang.org/x/exp/slog"
)

// ToggleCompleted toggles the completed status of a task.
func ToggleCompleted(ctx context.Context, db taskToggler, userID string, taskID string) error {
	t, err := db.GetUserTask(ctx, userID, taskID)
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
			slog.Info("task not found", "taskID", taskID)

			tsk := &Task{
				UserID:    ID(userID),
				TaskID:    ID(taskID),
				Completed: true,
				Hidden:    false,
			}

			err = db.AddUserTask(ctx, tsk)
			if err != nil {
				return err
			}

			return nil
		} else {
			return err
		}
	}

	t.Completed = !t.Completed

	err = db.UpdateUserTask(ctx, &t)
	if err != nil {
		return err
	}

	return nil
}

type taskToggler interface {
	GetUserTask(ctx context.Context, userID string, taskID string) (Task, error)
	AddUserTask(ctx context.Context, t *Task) error
	UpdateUserTask(ctx context.Context, t *Task) error
}
