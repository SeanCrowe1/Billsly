-- name: CreateTransaction :one
INSERT INTO transactions (id, created_at, updated_at, name, type, amount, due_date, bank, user_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;
