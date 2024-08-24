-- +goose Up
-- +goose StatementBegin
CREATE TYPE UsersStatus AS ENUM ('ACTIVE', 'INACTIVE');

CREATE TABLE IF NOT EXISTS users (
  "id" uuid PRIMARY KEY,
  "name" varchar(255) NOT NULL,
  "email" varchar(255) NOT NULL,
  "password" varchar(255) NOT NULL,
  "status" UsersStatus DEFAULT 'ACTIVE',
  "emailVerifiedAt" timestamp,
  "deletedAt" timestamp,
  "createdAt" timestamp NOT NULL,
  "updatedAt" timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TYPE IF EXISTS UsersStatus CASCADE;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
