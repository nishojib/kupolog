package task

import (
	"context"
	"errors"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
)

// ToggleCompleted toggles the completed status of a task.
func ToggleCompleted(
	ctx context.Context,
	db TaskToggler,
	userID string,
	taskID string,
	kind string,
) error {
	t, err := db.GetUserTask(ctx, userID, taskID)
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
			tsk := &Task{
				UserID:    ID(userID),
				TaskID:    ID(taskID),
				Completed: true,
				Hidden:    false,
				Kind:      Kind(kind),
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

// TaskToggler is an interface that represents the db operations for toggling tasks.

//go:generate mockery --with-expecter --name TaskToggler
type TaskToggler interface {
	GetUserTask(ctx context.Context, userID string, taskID string) (Task, error)
	AddUserTask(ctx context.Context, t *Task) error
	UpdateUserTask(ctx context.Context, t *Task) error
}
