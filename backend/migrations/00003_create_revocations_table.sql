-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS revocations (
  id INTEGER PRIMARY KEY,
  token TEXT KEY NOT NULL,
  revoked_at INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS revocations;
-- +goose StatementEnd
