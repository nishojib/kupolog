package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/auth"
	"github.com/nishojib/ffxivdailies/internal/task"
	"github.com/nishojib/ffxivdailies/internal/user"
)

// SharedTaskResponse represents the response for the shared tasks endpoint.
type SharedTaskResponse struct {
	Weeklies []TaskResponse `json:"weeklies,omitempty"`
	Dailies  []TaskResponse `json:"dailies,omitempty"`
}

// TaskResponse describes a task.
type TaskResponse struct {
	TaskID    string `json:"taskID"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
	Hidden    bool   `json:"hidden"`
}

// SharedTasksHandler godoc
//
//	@Summary		Shared tasks
//	@Description	Get the shared tasks
//	@Tags			tasks
//	@Produce		json
//	@Param			kind	query		string	true	"Kind of tasks to return"
//	@Success		200		{object}	SharedTaskResponse
//	@Failure		500		{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/tasks/shared [get]
func (s *Server) SharedTasksHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(user.ID)
	if userID == "" {
		api.AuthenticationRequiredResponse(w, r)
		return
	}

	input, err := task.GetTasks(r.Context(), s.db, string(userID))
	if err != nil {
		api.ServerErrorResponse(w, r, err)
		return
	}

	slog.Info("weeklies", "len", len(input.Weeklies))
	slog.Info("dailies", "len", len(input.Dailies))

	kind := r.URL.Query().Get("kind")

	var tasks []TaskResponse
	if kind == "weekly" {
		for _, task := range input.Weeklies {
			tasks = append(tasks, TaskResponse{
				TaskID:    string(task.TaskID),
				Title:     task.Title,
				Completed: bool(task.Completed),
				Hidden:    bool(task.Hidden),
			})
		}
		err = api.WriteJSON(w, http.StatusOK, SharedTaskResponse{Weeklies: tasks}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
			return
		}
	} else if kind == "daily" {
		for _, task := range input.Dailies {
			tasks = append(tasks, TaskResponse{
				TaskID:    string(task.TaskID),
				Title:     task.Title,
				Completed: bool(task.Completed),
				Hidden:    bool(task.Hidden),
			})
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
			weeklyTasks = append(weeklyTasks, TaskResponse{
				TaskID:    string(task.TaskID),
				Title:     task.Title,
				Completed: bool(task.Completed),
				Hidden:    bool(task.Hidden),
			})
		}

		for _, task := range input.Dailies {
			dailyTasks = append(dailyTasks, TaskResponse{
				TaskID:    string(task.TaskID),
				Title:     task.Title,
				Completed: bool(task.Completed),
				Hidden:    bool(task.Hidden),
			})
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

type ToggleTaskRequest struct {
	HasCompleted bool `json:"hasCompleted"`
	HasHidden    bool `json:"hasHidden"`
}

// ToogleTaskHandler godoc
//
//	@Summary		Toggle Task
//	@Description	toggle a task of the current user
//	@Tags			tasks
//	@Produce		json
//	@Param			request	body	ToggleTaskRequest	true	"request body"
//	@Param			taskID	path	string				true	"Task ID"
//	@Success		200
//	@Failure		400	{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500	{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/tasks/shared/{taskID} [put]
func (s *Server) ToggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value(auth.UserIDKey).(user.ID)
	if userID == "" {
		api.AuthenticationRequiredResponse(w, r)
		return
	}

	slog.Info("toggling task", "userID", userID)

	var input ToggleTaskRequest
	err := api.ReadJSON(w, r, &input)
	if err != nil {
		api.BadRequestResponse(w, r, err)
		return
	}

	taskId := chi.URLParam(r, "taskID")

	if input.HasCompleted {
		err = task.ToggleCompleted(r.Context(), s.db, string(userID), taskId)
	} else if input.HasHidden {
		slog.Info("toggling hidden", "taskID", taskId)
	}

	if err != nil {
		api.ServerErrorResponse(w, r, err)
		return
	}
}
