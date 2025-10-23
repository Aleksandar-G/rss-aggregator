-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: DeleteFeed :exec
DELETE FROM feeds WHERE ID=?;

-- name: GetFeedByName :one
SELECT * FROM feeds
WHERE name = ? LIMIT 1;

-- name: GetFeedById :one
SELECT * FROM feeds
WHERE ID = ? LIMIT 1;

-- name: ListFeeds :many
SELECT * FROM feeds;

-- name: UpdateFeed :exec
UPDATE feeds
SET name = ?
WHERE id = ?;