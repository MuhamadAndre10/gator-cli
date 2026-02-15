-- name: CreateFeedUser :one
WITH insert_feed_user AS (
    INSERT INTO feed_follow(id, user_id, feed_id)
    VALUES($1, $2, $3)
    RETURNING *
)
SELECT ifu.*,
    u.name AS user_name,
    f.name AS feed_name
FROM insert_feed_user ifu
    JOIN users u ON ifu.user_id = u.id
    JOIN feed f ON ifu.feed_id = f.id;
-- name: GetFollowingFeeds :many
SELECT f.name AS feed_name,
    f.url AS feed_url,
    u.name AS user_name
FROM feed_follow ff
    JOIN users u ON ff.user_id = u.id
    JOIN feed f ON ff.feed_id = f.id
WHERE ff.user_id = $1;
-- name: DeleteFeedUser :exec
DELETE FROM feed_follow
WHERE user_id = $1
    AND feed_id = $2;