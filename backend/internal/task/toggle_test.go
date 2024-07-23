package task_test

import (
	"context"
	"errors"
	"testing"

	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/task"
	"github.com/nishojib/ffxivdailies/internal/task/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestToggleCompleted_AddTask(t *testing.T) {
	dbError := errors.New("error")

	tests := map[string]struct {
		task              task.Task
		db                func(toggler *mocks.TaskToggler, tsk *task.Task)
		expectedErr       error
		expectedCompleted bool
	}{
		"toggle task when nothing exists": {
			task: task.Task{
				UserID:    task.ID("user-1234"),
				TaskID:    task.ID("task-1234"),
				Completed: false,
				Hidden:    false,
				Kind:      "daily",
			},
			db: func(toggler *mocks.TaskToggler, _ *task.Task) {
				toggler.EXPECT().
					GetUserTask(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(task.Task{}, repoErrors.ErrRecordNotFound)

				toggler.EXPECT().
					AddUserTask(mock.Anything, mock.AnythingOfType("*task.Task")).
					Return(nil)
			},
			expectedErr:       nil,
			expectedCompleted: false,
		},
		"error when adding task": {
			task: task.Task{},
			db: func(toggler *mocks.TaskToggler, _ *task.Task) {
				toggler.EXPECT().
					GetUserTask(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(task.Task{}, repoErrors.ErrRecordNotFound)

				toggler.EXPECT().
					AddUserTask(mock.Anything, mock.AnythingOfType("*task.Task")).
					Return(dbError)
			},
			expectedErr:       dbError,
			expectedCompleted: false,
		},
		"error when getting task": {
			task: task.Task{},
			db: func(toggler *mocks.TaskToggler, _ *task.Task) {
				toggler.EXPECT().
					GetUserTask(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(task.Task{}, dbError)

			},
			expectedErr:       dbError,
			expectedCompleted: false,
		},
		"toggle task when task exists": {
			task: task.Task{
				UserID:    task.ID("user-1234"),
				TaskID:    task.ID("task-1234"),
				Completed: false,
				Hidden:    false,
				Kind:      "daily",
			},
			db: func(toggler *mocks.TaskToggler, tsk *task.Task) {
				toggler.EXPECT().
					GetUserTask(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(*tsk, nil)

				toggler.EXPECT().
					UpdateUserTask(mock.Anything, mock.AnythingOfType("*task.Task")).
					RunAndReturn(func(_ context.Context, _ *task.Task) error {
						tsk.Completed = true
						return nil
					})
			},
			expectedErr:       nil,
			expectedCompleted: true,
		},
		"error when updating task": {
			task: task.Task{},
			db: func(toggler *mocks.TaskToggler, tsk *task.Task) {
				toggler.EXPECT().
					GetUserTask(mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("string")).
					Return(*tsk, nil)

				toggler.EXPECT().
					UpdateUserTask(mock.Anything, mock.AnythingOfType("*task.Task")).
					Return(dbError)
			},
			expectedErr:       dbError,
			expectedCompleted: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			taskToggler := mocks.NewTaskToggler(t)

			tc.db(taskToggler, &tc.task)

			err := task.ToggleCompleted(
				context.Background(),
				taskToggler,
				string(tc.task.TaskID),
				string(tc.task.UserID),
				string(tc.task.Kind),
			)

			require.ErrorIs(t, err, tc.expectedErr)

			if tc.expectedErr == nil {
				assert.Equal(t, tc.expectedCompleted, bool(tc.task.Completed))
			}
		})
	}
}
