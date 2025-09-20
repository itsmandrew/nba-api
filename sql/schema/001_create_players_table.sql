-- +goose Up
-- +goose StatementBegin
CREATE TABLE players (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  year_start INTEGER,
  year_end INTEGER,
  position VARCHAR(50),
  height VARCHAR(20),
  weight INTEGER,
  birth_date DATE,
  college VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS players;
-- +goose StatementEnd
