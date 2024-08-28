-- name: GetUser :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;
-- name: ListUsers :many
SELECT *
FROM users
ORDER BY card;
-- name: CreateUser :one
INSERT INTO users (email, card, verified)
VALUES ($1, $2, $3)
RETURNING *;
-- name: UpdateUser :exec
UPDATE users
set email = $2,
  card = $3,
  verified = $4
WHERE id = $1
RETURNING *;
-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;