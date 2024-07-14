package task

import (
	"context"
	"embed"
	"encoding/json"
	"slices"
	"strings"
)

//go:embed tasks.json
var f embed.FS

type Tasks struct {
	Weeklies []TaskResponse
	Dailies  []TaskResponse
}

type TaskResponse struct {
	TaskID    ID
	Title     string
	Completed Completed
	Hidden    Hidden
}

// GetTasks returns the the list of tasks
func GetTasks(ctx context.Context, db taskGetter, userID string) (Tasks, error) {
	var input struct {
		Weeklies []struct {
			TaskID string `json:"taskID"`
			Title  string `json:"title"`
		} `json:"weeklies"`
		Dailies []struct {
			TaskID string `json:"taskID"`
			Title  string `json:"title"`
		} `json:"dailies"`
	}

	buf, err := f.ReadFile("tasks.json")
	if err != nil {
		return Tasks{}, err
	}

	err = json.Unmarshal(buf, &input)
	if err != nil {
		return Tasks{}, err
	}

	ts, err := db.GetTasksForUser(ctx, userID)
	if err != nil {
		return Tasks{}, err
	}

	var tasks Tasks
	tasks.Weeklies = MergeTasks(input.Weeklies, ts)
	tasks.Dailies = MergeTasks(input.Dailies, ts)

	return tasks, nil
}

type taskGetter interface {
	GetTasksForUser(ctx context.Context, userID string) ([]Task, error)
}

func MergeTasks(t1 []struct {
	TaskID string `json:"taskID"`
	Title  string `json:"title"`
}, t2 []Task) []TaskResponse {
	taskMap := make(map[string]TaskResponse)

	for _, t := range t1 {
		taskMap[string(t.TaskID)] = TaskResponse{
			TaskID:    ID(t.TaskID),
			Title:     t.Title,
			Completed: false,
			Hidden:    false,
		}
	}

	for _, item := range t2 {
		if existing, ok := taskMap[string(item.TaskID)]; ok {
			existing.Completed = item.Completed
			existing.Hidden = item.Hidden
			taskMap[string(item.TaskID)] = existing
		}
	}

	var mergedTasks []TaskResponse
	for _, item := range taskMap {
		mergedTasks = append(mergedTasks, item)
	}

	slices.SortFunc(mergedTasks, func(a, b TaskResponse) int {
		return strings.Compare(string(a.TaskID), string(b.TaskID))
	})

	return mergedTasks
}
