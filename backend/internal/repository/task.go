package repository

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/task"
	"github.com/uptrace/bun"
)

// AddTaskToUser adds a task to a user.
func (r *Repository) AddUserTask(
	ctx context.Context,
	userTask *task.Task,
) error {
	_, err := r.db.NewInsert().Model(userTask).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUserTask updates a task for a user.
func (r *Repository) UpdateUserTask(ctx context.Context, t *task.Task) error {
	result, err := r.db.NewUpdate().
		Model(t).
		Set("completed = ?", t.Completed).
		Set("hidden = ?", t.Hidden).
		Set("version = ?", t.Version+1).
		Where("ut.user_id = ?", t.UserID).
		Where("ut.task_id = ?", t.TaskID).
		Where("ut.version = ?", t.Version).
		Exec(ctx)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repoErrors.ErrEditConflict
	}

	return nil
}

// GetUserTask returns a task for a user.
func (r *Repository) GetUserTask(
	ctx context.Context,
	userID string,
	taskID string,
) (task.Task, error) {
	var t task.Task
	err := r.db.NewSelect().
		Model(&t).
		Where("ut.user_id = ?", userID).
		Where("ut.task_id = ?", taskID).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return task.Task{}, repoErrors.ErrRecordNotFound
		default:
			return task.Task{}, err
		}
	}

	return t, nil
}

// GetTasksForUser returns the tasks for a user.
func (r *Repository) GetTasksForUser(ctx context.Context, userID string) ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.NewSelect().Model(&tasks).Where("ut.user_id = ?", userID).Scan(ctx)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []task.Task{}, nil
		}

		return nil, err
	}

	return tasks, nil
}

func (r *Repository) UpdateTaskForKind(ctx context.Context, kind string) error {
	err := r.db.RunInTx(ctx, nil, func(innerCtx context.Context, tx bun.Tx) error {
		var ts []task.Task
		err := tx.NewSelect().Model(&ts).Where("kind = ?", kind).Scan(innerCtx)
		if err != nil {
			slog.Error("failed to get tasks", "error", err)
			return err
		}

		for _, t := range ts {
			if !t.Completed {
				continue
			}

			result, err := tx.NewUpdate().
				Model(&t).
				Set("completed = ?", false).
				Set("version = ?", t.Version+1).
				Where("ut.user_id = ?", t.UserID).
				Where("ut.task_id = ?", t.TaskID).
				Where("ut.version = ?", t.Version).
				Exec(innerCtx)

			if err != nil {
				slog.Error("failed to update task", "error", err)
				return err
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				slog.Error("failed to get rows affected", "error", err)
				return err
			}

			if rowsAffected == 0 {
				slog.Error("task not found", "taskID", t.TaskID)
				return err
			}
		}

		return nil
	})

	return err
}
