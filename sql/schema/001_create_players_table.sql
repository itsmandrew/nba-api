-- +goose Up
-- +goose StatementBegin
CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  year_start INTEGER NOT NULL DEFAULT 0,
  year_end INTEGER NOT NULL DEFAULT 0,
  position VARCHAR(50) NOT NULL DEFAULT '',
  height VARCHAR(20) NOT NULL DEFAULT '',
  weight INTEGER NOT NULL DEFAULT 0,
  birth_date DATE NOT NULL DEFAULT '1900-01-01',
  college VARCHAR(255) NOT NULL DEFAULT '',
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS players;
-- +goose StatementEnd
