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


-- name: GetFeedByURL :one
SELECT * FROM feeds WHERE feed_url = $1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = CURRENT_TIMESTAMP, updated_at = CURRENT_TIMESTAMP
WHERE id = $1;

-- name: GetNextFeedToFetch :one
SELECT * FROM feeds
ORDER BY last_fetched_at NULLS FIRST
LIMIT 1;