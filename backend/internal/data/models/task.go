package models

import (
	"database/sql"
	"time"

	"github.com/nishojib/ffxivdailies/internal/validator"
)

// Task represents a task in the database.
type Task struct {
	ID          int64          `json:"-"                  bun:"id,pk,autoincrement"`
	TaskID      string         `json:"taskID"`
	Title       string         `json:"title"`
	Completed   bool           `json:"completed"`
	ContentType string         `json:"contentType"`
	Kind        string         `json:"kind"`
	IsHidden    bool           `json:"isHidden"`
	CreatedAt   time.Time      `json:"createdAt"          bun:",default:current_timestamp"`
	Version     int64          `json:"-"`
	CreatorID   sql.NullString `json:"-"`
	Subtasks    []*Subtask     `json:"subtasks,omitempty" bun:"rel:has-many,join:task_id=task_id"`
}

func (t *Task) Validate(v *validator.Validator) {
	v.Check(t.TaskID != "", "task_id", "must be provided")

	v.Check(t.Title != "", "title", "must be provided")

	v.Check(t.Kind != "", "kind", "must be provided")
	v.Check(t.Kind == "daily" || t.Kind == "weekly", "kind", "must be either daily or weekly")

	v.Check(t.ContentType != "", "type", "must be provided")
	v.Check(
		t.ContentType == "none" || t.ContentType == "subtask" || t.ContentType == "image" ||
			t.ContentType == "embed",
		"type",
		"must be either subtask, image or embed",
	)

	v.Check(t.CreatorID.Valid && t.CreatorID.String != "", "creator_id", "must be provided")
}
