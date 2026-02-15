-- +goose up
CREATE TABLE posts (
    ID uuid PRIMARY KEY,
    -- fk key
    feed_id uuid NOT NULL,
    -- post data
    title TEXT NOT NULL,
    url TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP,
    -- timestamp
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    -- add constraint
    CONSTRAINT fk_feed FOREIGN KEY (feed_id) REFERENCES feed(ID) ON DELETE CASCADE
);
-- +goose down
DROP TABLE IF EXISTS posts;