package server

import (
	"embed"
	"encoding/json"
	"net/http"

	"github.com/nishojib/ffxivdailies/internal/api"
)

//go:embed tasks.json
var f embed.FS

// SharedTaskResponse represents the response for the shared tasks endpoint.
type SharedTaskResponse struct {
	Weeklies []TaskResponse `json:"weeklies,omitempty"`
	Dailies  []TaskResponse `json:"dailies,omitempty"`
}

// TaskResponse describes a task.
type TaskResponse struct {
	TaskID string `json:"taskID"`
	Title  string `json:"title"`
}

// Weeklies godoc
//
//	@Summary		Shared tasks
//	@Description	Get the shared tasks
//	@Tags			tasks
//	@Produce		json
//	@Param			kind	query		string	true	"Kind of tasks to return"
//	@Success		200	{object}	SharedTaskResponse
//	@Failure		500	{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/tasks/shared [get]
func (s *Server) SharedTasksHandler(w http.ResponseWriter, r *http.Request) {
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
		api.ServerErrorResponse(w, r, err)
		return
	}

	err = json.Unmarshal(buf, &input)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
		return
	}

	kind := r.URL.Query().Get("kind")

	var tasks []TaskResponse
	if kind == "weekly" {
		for _, task := range input.Weeklies {
			tasks = append(tasks, TaskResponse{TaskID: task.TaskID, Title: task.Title})
		}
		err = api.WriteJSON(w, http.StatusOK, SharedTaskResponse{Weeklies: tasks}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	} else if kind == "daily" {
		for _, task := range input.Dailies {
			tasks = append(tasks, TaskResponse{TaskID: task.TaskID, Title: task.Title})
		}

		err = api.WriteJSON(w, http.StatusOK, SharedTaskResponse{Dailies: tasks}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	} else {
		weeklyTasks := []TaskResponse{}
		dailyTasks := []TaskResponse{}

		for _, task := range input.Weeklies {
			weeklyTasks = append(weeklyTasks, TaskResponse{TaskID: task.TaskID, Title: task.Title})
		}

		for _, task := range input.Dailies {
			dailyTasks = append(dailyTasks, TaskResponse{TaskID: task.TaskID, Title: task.Title})
		}
		err = api.WriteJSON(w, http.StatusOK, SharedTaskResponse{
			Weeklies: weeklyTasks,
			Dailies:  dailyTasks,
		}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	}
}
