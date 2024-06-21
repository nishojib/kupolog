-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS accounts (
  id INTEGER PRIMARY KEY,
  provider TEXT NOT NULL,
  provider_account_id TEXT NOT NULL,
  email TEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at TIMESTAMP,
  version INTEGER NOT NULL DEFAULT 1,

  user_id INTEGER NOT NULL REFERENCES users (id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX IF NOT EXISTS accounts_provider_account_id_index ON accounts(provider_account_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS accounts_provider_account_id_index;

DROP TABLE IF EXISTS accounts;
-- +goose StatementEnd
