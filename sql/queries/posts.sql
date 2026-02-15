-- name: CreatePosts :one
INSERT INTO posts (
        ID,
        feed_id,
        title,
        url,
        description,
        published_at
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetPostsForUser :many
SELECT posts.*,
    feed.name AS feed_name
FROM posts
    JOIN feed_follow ON feed_follow.feed_id = posts.feed_id
    JOIN feed ON posts.feed_id = feed.id
WHERE feed_follow.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2;