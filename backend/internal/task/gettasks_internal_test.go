package task

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeTasks(t *testing.T) {
	t1 := []FileTask{
		{TaskID: "task1", Title: "Task 1"},
		{TaskID: "task2", Title: "Task 2"},
		{TaskID: "task3", Title: "Task 3"},
	}

	t2 := []Task{
		{TaskID: "task1", Completed: true},
		{TaskID: "task2", Completed: false},
	}

	mergedTasks := mergeTasks(t1, t2)

	expected := []TaskResponse{
		{TaskID: "task1", Title: "Task 1", Completed: true, Hidden: false},
		{TaskID: "task2", Title: "Task 2", Completed: false, Hidden: false},
		{TaskID: "task3", Title: "Task 3", Completed: false, Hidden: false},
	}

	assert.Equal(t, expected, mergedTasks)
	assert.Len(t, mergedTasks, 3)
}
