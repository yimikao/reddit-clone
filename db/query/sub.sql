-- name: CreateSub :one
INSERT INTO subs (
    name,
    creator_id
) VALUES (
    $1, $2
) RETURNING *;


-- name: GetSub :one 
SELECT * FROM subs
WHERE name = $1 LIMIT 1;

-- name: ListSub :many
SELECT * FROM subs
ORDER BY name
LIMIT $1
OFFSET $2; 

-- name: UpdateSub :one
UPDATE subs
SET 
    name = $2,
    avatar = $3,
    description = $4
WHERE name = $1
RETURNING *;

-- name: DeleteSub :exec
DELETE FROM subs 
WHERE name = $1;