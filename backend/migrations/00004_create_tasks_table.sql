-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS tasks (
  id INTEGER PRIMARY KEY,
  task_id TEXT NOT NULL UNIQUE,
  title TEXT NOT NULL UNIQUE,
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  content_type TEXT NOT NULL,
  kind TEXT NOT NULL,
  is_hidden BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  version INTEGER NOT NULL DEFAULT 1,

  creator_id TEXT DEFAULT NULL REFERENCES users(user_id) ON DELETE CASCADE
);


CREATE UNIQUE INDEX IF NOT EXISTS tasks_task_id_index ON tasks(task_id);
CREATE INDEX IF NOT EXISTS tasks_creator_id_index ON tasks(creator_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS tasks_task_id_index;
DROP INDEX IF EXISTS tasks_creator_id_index;

DROP TABLE IF EXISTS tasks;
-- +goose StatementEnd
