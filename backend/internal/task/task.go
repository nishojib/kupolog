package task

import "github.com/uptrace/bun"

// ID is an identifier.
type ID string

// Completed represents whether a task is completed.
type Completed bool

// Hidden represents whether a task is hidden.
type Hidden bool

// Kind represents the kind of task (daily | weekly | custom).
type Kind string

// Task represents a task for a user.
type Task struct {
	bun.BaseModel `bun:"table:users_tasks,alias:ut"`

	TaskID    ID        `bun:"task_id"`
	UserID    ID        `bun:"user_id"`
	Completed Completed `bun:"completed"`
	Hidden    Hidden    `bun:"hidden"`
	Kind      Kind      `bun:"kind"`
	Version   int64
}
