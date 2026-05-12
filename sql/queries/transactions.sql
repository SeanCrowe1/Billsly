-- name: CreateTransaction :one
INSERT INTO transactions (id, created_at, updated_at, t_name, t_type, amount, due_date, bank, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetTransactionByName :one
SELECT * FROM transactions
WHERE t_name = $1;

-- name: GetTransactionsForUser :many
SELECT * FROM transactions
Where user_id = $1;

-- name: GetInTransactionsForUser :many
SELECT * FROM transactions
Where user_id = $1
AND t_type = 'in'
ORDER BY due_date, t_name;

-- name: GetOutTransactionsForUser :many
SELECT * FROM transactions
Where user_id = $1
AND t_type = 'out'
ORDER BY due_date, t_name;

-- name: GetAllBanks :many
SELECT DISTINCT bank FROM transactions;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;
