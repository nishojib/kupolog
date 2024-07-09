package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/nishojib/ffxivdailies/internal/data/models"

	"github.com/uptrace/bun"
)

type TaskRepository struct {
	db *bun.DB
}

func NewTaskRepository(db *bun.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (tr *TaskRepository) GetByKind(kind string) ([]models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var tasks []models.Task
	err := tr.db.NewSelect().
		Model(&tasks).
		Relation("Subtasks").
		Where("kind = ?", kind).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return []models.Task{}, ErrRecordNotFound
		default:
			return []models.Task{}, err
		}
	}

	return tasks, nil
}

func (tr *TaskRepository) GetByID(taskID string) (models.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var task models.Task
	err := tr.db.NewSelect().
		Model(&task).
		Relation("Subtasks").
		Where("task_id = ?", taskID).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.Task{}, ErrRecordNotFound
		default:
			return models.Task{}, err
		}
	}

	return task, nil
}

func (tr *TaskRepository) ToggleTask(task *models.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := tr.db.RunInTx(ctx, &sql.TxOptions{}, func(ctx context.Context, tx bun.Tx) error {
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
				return ErrEditConflict
			default:
				return err
			}
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if rowsAffected == 0 {
			return ErrEditConflict
		}

		for _, subtask := range task.Subtasks {
			result, err = tx.NewUpdate().
				Model(subtask).
				Set("completed = ?", !task.Completed).
				Set("version = ?", subtask.Version+1).
				Where("subtask_id = ?", subtask.SubtaskID).
				Where("version = ?", subtask.Version).
				Exec(ctx)
			if err != nil {
				switch {
				case errors.Is(err, sql.ErrNoRows):
					return ErrEditConflict
				default:
					return err
				}
			}

			rowsAffected, err := result.RowsAffected()
			if err != nil {
				return err
			}

			if rowsAffected == 0 {
				return ErrEditConflict
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (tr *TaskRepository) GetSubtaskByID(subtaskID string) (models.Subtask, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var subtask models.Subtask
	err := tr.db.NewSelect().
		Model(&subtask).
		Where("subtask_id = ?", subtaskID).
		Scan(ctx)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return models.Subtask{}, ErrRecordNotFound
		default:
			return models.Subtask{}, err
		}
	}

	return subtask, nil
}

func (tr *TaskRepository) ToggleSubtask(subtask *models.Subtask) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := tr.db.NewUpdate().
		Model(subtask).
		Set("completed = ?", !subtask.Completed).
		Set("version = ?", subtask.Version+1).
		Where("subtask_id = ?", subtask.SubtaskID).
		Where("version = ?", subtask.Version).
		Exec(ctx)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrEditConflict
	}

	return nil
}
