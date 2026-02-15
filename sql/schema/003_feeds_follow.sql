-- +goose Up
CREATE TABLE IF NOT EXISTS feed_follow (
    id uuid PRIMARY KEY,
    user_id uuid NOT NULL,
    feed_id uuid NOT NULL,
    -- timestamps
    created_at timestamp NOT NULL DEFAULT NOW(),
    updated_at timestamp NOT NULL DEFAULT NOW(),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (feed_id) REFERENCES feed(id) ON DELETE CASCADE,
    UNIQUE (user_id, feed_id)
);
CREATE INDEX idx_feed_follow_user_id ON feed_follow(user_id);
CREATE INDEX idx_feed_follow_feed_id ON feed_follow(feed_id);
-- +goose Down
DROP INDEX IF EXISTS idx_feed_follow_user_id;
DROP INDEX IF EXISTS idx_feed_follow_feed_id;
DROP TABLE IF EXISTS feed_follow;