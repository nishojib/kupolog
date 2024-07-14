-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_tasks (
  user_id TEXT NOT NULL REFERENCES users (user_id) ON DELETE CASCADE,
  task_id TEXT NOT NULL,
  version INTEGER NOT NULL DEFAULT 1,
  completed BOOLEAN NOT NULL DEFAULT FALSE,
  hidden BOOLEAN NOT NULL DEFAULT FALSE,
  PRIMARY KEY (user_id, task_id)
);

CREATE INDEX IF NOT EXISTS users_tasks_user_id_index ON users_tasks(user_id);
CREATE INDEX IF NOT EXISTS users_tasks_task_id_index ON users_tasks(task_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS users_tasks_user_id_index;
DROP INDEX IF EXISTS users_tasks_task_id_index;

DROP TABLE IF EXISTS users_tasks;
-- +goose StatementEnd
