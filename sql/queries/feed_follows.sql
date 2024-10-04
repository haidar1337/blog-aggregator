-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, user_id, feed_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)

SELECT 
    inserted_feed_follow.*,
    users.name AS user_name,
    feeds.feed_name AS feed_name
FROM inserted_feed_follow
INNER JOIN users
ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds
ON feeds.id = inserted_feed_follow.feed_id;

-- name: GetFeedFollowsForUser :many
SELECT feed_name, feed_url 
FROM feed_follows 
INNER JOIN feeds
ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1;