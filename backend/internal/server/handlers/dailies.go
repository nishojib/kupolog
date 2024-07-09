package handlers

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nishojib/ffxivdailies/internal/api"
	"github.com/nishojib/ffxivdailies/internal/data/models"
	"github.com/nishojib/ffxivdailies/internal/data/repository"
	"github.com/uptrace/bun"
)

// Weeklies godoc
//
//	@Summary		Weekly tasks
//	@Description	Get the weekly tasks
//	@Tags			dailies
//	@Produce		json
//	@Success		200	{object}	object{weeklies=[]models.Task}
//	@Failure		404	{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500	{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/dailies/weekly [get]
func Weeklies(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := repository.NewTaskRepository(db).GetByKind("weekly")
		if err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				api.NotFoundResponse(w, r)
				return
			}

			api.ServerErrorResponse(w, r, err)
			return
		}

		if len(data) == 0 {
			api.NotFoundResponse(w, r)
			return
		}

		err = api.WriteJSON(w, http.StatusOK, api.Envelope[[]models.Task]{"weeklies": data}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
		}
	}
}

// Dailies godoc
//
//	@Summary		Daily tasks
//	@Description	Get the daily tasks
//	@Tags			dailies
//	@Produce		json
//	@Success		200	{object}	object{dailies=[]models.Task}
//	@Failure		404	{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500	{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/dailies/daily [get]
func Dailies(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := repository.NewTaskRepository(db).GetByKind("daily")
		if err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				api.NotFoundResponse(w, r)
				return
			}

			api.ServerErrorResponse(w, r, err)
			return
		}

		if len(data) == 0 {
			api.NotFoundResponse(w, r)
			return
		}

		err = api.WriteJSON(w, http.StatusOK, api.Envelope[[]models.Task]{"dailies": data}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
		}
	}
}

// ToggleTask godoc
//
//	@Summary		Toggle task
//	@Description	Toggle the task
//	@Tags			dailies
//	@Accept			json
//	@Produce		json
//	@Param			taskID	path		string	true	"Task ID"
//	@Success		200		{object}	object{task=models.Task}
//	@Failure		404		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500		{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/dailies/tasks/{taskID} [put]
func ToggleTask(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		taskID := chi.URLParam(r, "taskID")
		task, err := repository.NewTaskRepository(db).GetByID(taskID)
		if err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				api.NotFoundResponse(w, r)
				return
			}

			api.ServerErrorResponse(w, r, err)
			return
		}

		err = repository.NewTaskRepository(db).ToggleTask(&task)
		if err != nil {
			if errors.Is(err, repository.ErrEditConflict) {
				api.EditConflictResponse(w, r)
				return
			}

			api.ServerErrorResponse(w, r, err)
			return
		}

		err = api.WriteJSON(w, http.StatusOK, api.Envelope[models.Task]{"task": task}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
		}
	}
}

// ToggleSubtask godoc
//
//	@Summary		Toggle subtask
//	@Description	Toggle the subtask
//	@Tags			dailies
//	@Accept			json
//	@Produce		json
//	@Param			subtaskID	path		string	true	"Subtask ID"
//	@Success		200		{object}	object{task=models.Subtask}
//	@Failure		404		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500		{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/dailies/subtasks/{subtaskID} [put]
func ToggleSubtask(db *bun.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		subtaskID := chi.URLParam(r, "subtaskID")
		subtask, err := repository.NewTaskRepository(db).GetSubtaskByID(subtaskID)
		if err != nil {
			if errors.Is(err, repository.ErrRecordNotFound) {
				api.NotFoundResponse(w, r)
				return
			}

			api.ServerErrorResponse(w, r, err)
			return
		}

		err = repository.NewTaskRepository(db).ToggleSubtask(&subtask)
		if err != nil {
			if errors.Is(err, repository.ErrEditConflict) {
				api.EditConflictResponse(w, r)
				return
			}

			api.ServerErrorResponse(w, r, err)
			return
		}

		err = api.WriteJSON(w, http.StatusOK, api.Envelope[models.Subtask]{"subtask": subtask}, nil)
		if err != nil {
			api.ServerErrorResponse(w, r, err)
		}
	}
}
