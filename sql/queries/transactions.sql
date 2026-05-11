-- name: CreateTransaction :one
INSERT INTO transactions (id, created_at, updated_at, name, type, amount, due_date, bank, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetTransactionByName :one
SELECT * FROM transactions
WHERE name = $1;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;
