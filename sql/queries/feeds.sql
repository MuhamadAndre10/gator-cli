-- name: CreateFeed :one
INSERT INTO feed (id, name, url, user_id)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetFeedByName :one
SELECT *
FROM feed
WHERE name = $1;
-- name: DeleteAllFeed :exec
DELETE FROM feed;
-- name: GetAllFeeds :many
SELECT name,
    url,
    user_id
FROM feed;
-- name: GetUserFeeds :many
SELECT f.name AS feed_name,
    f.url AS feed_url,
    u.name AS user_name
FROM feed f
    INNER JOIN users u ON f.user_id = u.id
ORDER BY u.name;
-- name: GetFeedByUrl :one
SELECT *
FROM feed
WHERE url = $1;
-- name: MarkFeedFetched :one
UPDATE feed
SET last_fetched_at = NOW(),
    updated_at = NOW()
WHERE id = $1
RETURNING *;
-- name: GetNextFeedToFetch :one
SELECT *
FROM feed
ORDER BY last_fetched_at ASC NULLS FIRST
LIMIT 1;