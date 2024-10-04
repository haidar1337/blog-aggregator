// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: feed_follows.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createFeedFollow = `-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
    INSERT INTO feed_follows(id, user_id, feed_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING id, user_id, feed_id, created_at, updated_at
)

SELECT 
    inserted_feed_follow.id, inserted_feed_follow.user_id, inserted_feed_follow.feed_id, inserted_feed_follow.created_at, inserted_feed_follow.updated_at,
    users.name AS user_name,
    feeds.feed_name AS feed_name
FROM inserted_feed_follow
INNER JOIN users
ON users.id = inserted_feed_follow.user_id
INNER JOIN feeds
ON feeds.id = inserted_feed_follow.feed_id
`

type CreateFeedFollowParams struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CreateFeedFollowRow struct {
	ID        uuid.UUID
	UserID    uuid.UUID
	FeedID    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	UserName  string
	FeedName  string
}

func (q *Queries) CreateFeedFollow(ctx context.Context, arg CreateFeedFollowParams) (CreateFeedFollowRow, error) {
	row := q.db.QueryRowContext(ctx, createFeedFollow,
		arg.ID,
		arg.UserID,
		arg.FeedID,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i CreateFeedFollowRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.FeedID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.UserName,
		&i.FeedName,
	)
	return i, err
}

const getFeedFollowsForUser = `-- name: GetFeedFollowsForUser :many
SELECT feed_name, feed_url 
FROM feed_follows 
INNER JOIN feeds
ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
`

type GetFeedFollowsForUserRow struct {
	FeedName string
	FeedUrl  string
}

func (q *Queries) GetFeedFollowsForUser(ctx context.Context, userID uuid.UUID) ([]GetFeedFollowsForUserRow, error) {
	rows, err := q.db.QueryContext(ctx, getFeedFollowsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetFeedFollowsForUserRow
	for rows.Next() {
		var i GetFeedFollowsForUserRow
		if err := rows.Scan(&i.FeedName, &i.FeedUrl); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}