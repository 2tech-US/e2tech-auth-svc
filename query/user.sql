-- name: CreateUser :one
INSERT INTO users (
  phone,
  password,
  name,
  role
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: GetUserByPhone :one
SELECT * FROM users
WHERE phone = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2; -- pagination: offset: skip many rows

-- name: UpdatePassword :one
UPDATE users
SET password = $2
WHERE id = $1
RETURNING *;

-- name: Verify :one
UPDATE users
SET verified = true
WHERE phone = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;