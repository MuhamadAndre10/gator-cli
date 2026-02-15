-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id uuid PRIMARY KEY,
    name varchar(100) UNIQUE NOT NULL,
    -- timestamps
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL
);
-- +goose Down
DROP TABLE IF EXISTS users;