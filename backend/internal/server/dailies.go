package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/nishojib/ffxivdailies/internal/api"
	repoErrors "github.com/nishojib/ffxivdailies/internal/errors"
	"github.com/nishojib/ffxivdailies/internal/task"
)

// Weeklies godoc
//
//	@Summary		Weekly tasks
//	@Description	Get the weekly tasks
//	@Tags			dailies
//	@Produce		json
//	@Success		200	{object}	object{weeklies=[]task.Task}
//	@Failure		404	{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500	{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/dailies/weekly [get]
func (s *Server) WeekliesHandler(w http.ResponseWriter, r *http.Request) {
	dbCtx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	data, err := s.db.GetTasksByKind(dbCtx, "weekly")
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
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

	err = api.WriteJSON(w, http.StatusOK, api.Envelope[[]task.Task]{"weeklies": data}, nil)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
	}
}

// Dailies godoc
//
//	@Summary		Daily tasks
//	@Description	Get the daily tasks
//	@Tags			dailies
//	@Produce		json
//	@Success		200	{object}	object{dailies=[]task.Task}
//	@Failure		404	{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500	{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/dailies/daily [get]
func (s *Server) DailiesHandler(w http.ResponseWriter, r *http.Request) {
	dbCtx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	data, err := s.db.GetTasksByKind(dbCtx, "daily")
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
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

	err = api.WriteJSON(w, http.StatusOK, api.Envelope[[]task.Task]{"dailies": data}, nil)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
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
//	@Success		200		{object}	object{task=task.Task}
//	@Failure		404		{object}	object{detail=string,status=int,title=string,type=string}
//	@Failure		500		{object}	object{detail=string,status=int,title=string,type=string}
//	@Router			/dailies/tasks/{taskID} [put]
func (s *Server) ToggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	dbCtx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	taskID := chi.URLParam(r, "taskID")
	t, err := s.db.GetTaskByID(dbCtx, taskID)
	if err != nil {
		if errors.Is(err, repoErrors.ErrRecordNotFound) {
			api.NotFoundResponse(w, r)
			return
		}

		api.ServerErrorResponse(w, r, err)
		return
	}

	err = s.db.ToggleTask(dbCtx, &t)
	if err != nil {
		if errors.Is(err, repoErrors.ErrEditConflict) {
			api.EditConflictResponse(w, r)
			return
		}

		api.ServerErrorResponse(w, r, err)
		return
	}

	err = api.WriteJSON(w, http.StatusOK, api.Envelope[task.Task]{"task": t}, nil)
	if err != nil {
		api.ServerErrorResponse(w, r, err)
	}
}
