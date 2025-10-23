-- name: CreateUserFeed :one
INSERT INTO users_feeds (id, created_at, updated_at, user_id, feed_id)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteUserFeed :exec
DELETE FROM users_feeds WHERE ID=?;

-- name: GetUserFeedById :one
SELECT * FROM users_feeds
WHERE ID = ? LIMIT 1;

-- name: ListUserFeeds :many
SELECT * FROM users_feeds;
