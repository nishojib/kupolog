package models

import (
	"database/sql"
	"time"

	"github.com/nishojib/ffxivdailies/internal/validator"
)

// Subtask represents a subtask in the database.
type Subtask struct {
	ID        int64          `json:"-"         bun:"id,pk,autoincrement"`
	SubtaskID string         `json:"subtaskID"`
	Title     string         `json:"title"`
	Completed bool           `json:"completed"`
	IsHidden  bool           `json:"isHidden"`
	CreatedAt time.Time      `json:"createdAt" bun:",default:current_timestamp"`
	Version   int64          `json:"-"`
	TaskID    string         `json:"-"`
	CreatorID sql.NullString `json:"-"`
}

func (st *Subtask) Validate(v *validator.Validator) {
	v.Check(st.SubtaskID != "", "subtask_id", "must be provided")

	v.Check(st.Title != "", "title", "must be provided")

	v.Check(st.CreatorID.Valid && st.CreatorID.String != "", "creator_id", "must be provided")
}
