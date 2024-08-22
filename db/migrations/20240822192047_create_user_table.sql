-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "createdAt" timestamp NOT NULL,
  "updatedAt" timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users
-- +goose StatementEnd
