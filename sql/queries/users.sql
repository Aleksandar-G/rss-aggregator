-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, name)
VALUES (?, ?, ?, ?)
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE ID=?;

-- name: GetUserByName :one
SELECT * FROM users
WHERE name = ? LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE ID = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users;

-- name: UpdateAuthor :exec
UPDATE users
SET name = ?
WHERE id = ?;