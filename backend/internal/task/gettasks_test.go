package task_test

import (
	"context"
	"errors"
	"testing"

	"github.com/nishojib/ffxivdailies/internal/task"
	"github.com/nishojib/ffxivdailies/internal/task/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetTasks(t *testing.T) {
	taskGetter := mocks.NewTaskGetter(t)

	taskGetter.EXPECT().
		GetTasksForUser(mock.Anything, mock.AnythingOfType("string")).
		Return([]task.Task{
			{
				TaskID:    "weekly_repeatable_quests",
				UserID:    "user-1234",
				Completed: true,
				Hidden:    false,
				Kind:      "weekly",
			},
		}, nil)

	tasks, err := task.GetTasks(context.Background(), taskGetter, "user-1234")
	require.NoError(t, err)

	for _, task := range tasks.Weeklies {
		if task.TaskID != "weekly_repeatable_quests" {
			continue
		}
		assert.Equal(t, bool(task.Completed), true)
	}
}

func TestGetTasks_Error(t *testing.T) {
	dbError := errors.New("error")

	taskGetter := mocks.NewTaskGetter(t)

	taskGetter.EXPECT().
		GetTasksForUser(mock.Anything, mock.AnythingOfType("string")).
		Return([]task.Task{}, dbError)

	_, err := task.GetTasks(context.Background(), taskGetter, "user-1234")
	require.ErrorIs(t, err, dbError)
}
