-- name: CreateUser :one
INSERT INTO users (
  username,
  hashed_password,
  email
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET hashed_password = $2,
    avatar = $3,
    bio = $4
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: ChangeBlockedStatus :exec
UPDATE users
SET isBlocked = $2
WHERE id = $1;
