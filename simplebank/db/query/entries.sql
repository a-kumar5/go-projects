-- name: CreateEntry :one
INSERT INTO entries (
  account_id,
  amount
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: ListEntry :many
SELECT * FROM entries
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
set balance = $2
WHERE id = $1
RETURNING *;