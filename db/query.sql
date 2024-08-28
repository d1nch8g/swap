-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY name;
-- name: CreateUser :one
INSERT INTO users (email, card, verified)
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateUser :exec
UPDATE users
set name = $2,
  bio = $3
WHERE id = $1
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;