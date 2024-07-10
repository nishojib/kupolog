package repository

import (
	"context"
	"database/sql"
	"errors"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/task"

	"github.com/uptrace/bun"
)

func (r *Repository) GetTasksByKind(ctx context.Context, kind string) ([]task.Task, error) {
	var tasks []task.Task
	err := r.db.NewSelect().
		Model(&tasks).
		Relation("Subtasks").
		Where("kind = ?", kind).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return []task.Task{}, repoErrors.ErrRecordNotFound
		default:
			return []task.Task{}, err
		}
	}

	return tasks, nil
}

func (r *Repository) GetTaskByID(ctx context.Context, taskID string) (task.Task, error) {
	var t task.Task
	err := r.db.NewSelect().
		Model(&t).
		Relation("Subtasks").
		Where("task_id = ?", taskID).
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

func (r *Repository) ToggleTask(ctx context.Context, task *task.Task) error {
	err := r.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
		result, err := tx.NewUpdate().
			Model(task).
			Set("completed = ?", !task.Completed).
			Set("version = ?", task.Version+1).
			Where("id = ?", task.ID).
			Where("version = ?", task.Version).
			Exec(ctx)
		if err != nil {
			switch {
			case errors.Is(err, sql.ErrNoRows):
				return repoErrors.ErrEditConflict
			default:
				return err
			}
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return repoErrors.ErrEditConflict
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
