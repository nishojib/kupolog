package models

import (
	"time"

	"github.com/nishojib/ffxivdailies/internal/validator"
)

// User represents a user in the database.
type Task struct {
	ID        int64     `json:"-"          bun:"id,pk,autoincrement"`
	TaskID    string    `json:"task_id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	TaskType  string    `json:"type"`
	Kind      string    `json:"kind"`
	IsHidden  bool      `json:"is_hidden"`
	CreatedAt time.Time `json:"created_at" bun:",default:current_timestamp"`
	Version   int64     `json:"-"`
	CreatorID *string   `json:"-"`
	ParentID  *string   `json:"-"`
}

func (t *Task) Validate(v *validator.Validator) {
	v.Check(t.TaskID != "", "task_id", "must be provided")

	v.Check(t.Title != "", "title", "must be provided")

	v.Check(t.Kind != "", "kind", "must be provided")
	v.Check(t.Kind == "daily" || t.Kind == "weekly", "kind", "must be either daily or weekly")

	v.Check(t.TaskType != "", "type", "must be provided")
	v.Check(
		t.TaskType == "subtask" || t.TaskType == "image" || t.TaskType == "embed",
		"type",
		"must be either subtask, image or embed",
	)

	v.Check(*t.CreatorID != "", "creator_id", "must be provided")
}
