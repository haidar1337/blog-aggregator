-- name: CreateFeed :one
INSERT INTO feeds(id, feed_name, feed_url, user_id, created_at, updated_at)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)

RETURNING *;

-- name: GetFeedsWithUsers :many
SELECT feed_name, feed_url, users.name FROM feeds
INNER JOIN users
ON users.id = feeds.user_id;