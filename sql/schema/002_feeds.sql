-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id)
    ON DELETE CASCADE, 
    feed_name TEXT NOT NULL, 
    feed_url TEXT NOT NULL UNIQUE,

    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down
DROP TABLE feeds;