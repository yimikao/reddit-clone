-- name: CreateSubscriber :one
INSERT INTO subscribers (
    sub_id,
    user_id    
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetSubscriber :one
SELECT * FROM subscribers
WHERE sub_id = $1 AND user_id = $2
LIMIT 1;

-- name: DeleteSubscriber :exec
DELETE FROM subscribers
WHERE id = $1;