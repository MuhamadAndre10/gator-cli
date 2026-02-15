-- +goose Up
CREATE TABLE IF NOT EXISTS feed (
    id uuid PRIMARY KEY,
    name varchar(100) UNIQUE NOT NULL,
    url varchar(100) UNIQUE NOT NULL,
    user_id uuid NOT NULL,
    -- timestamps
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
-- +goose Down
DROP TABLE IF EXISTS feed;