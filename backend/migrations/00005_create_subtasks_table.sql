-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS subtasks (
  id INTEGER PRIMARY KEY,
  subtask_id TEXT NOT NULL UNIQUE,
  title TEXT NOT NULL UNIQUE,
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  is_hidden BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  version INTEGER NOT NULL DEFAULT 1,

  task_id INTEGER NOT NULL REFERENCES tasks(task_id) ON DELETE CASCADE,
  creator_id TEXT REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS subtasks_subtask_id_index ON subtasks(subtask_id);
CREATE INDEX IF NOT EXISTS subtasks_creator_id_index ON subtasks(creator_id);
CREATE INDEX IF NOT EXISTS subtasks_task_id_index ON subtasks(task_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS subtasks_subtask_id_index;
DROP INDEX IF EXISTS subtasks_creator_id_index;
DROP INDEX IF EXISTS subtasks_task_id_index;

DROP TABLE IF EXISTS subtasks;
-- +goose StatementEnd
