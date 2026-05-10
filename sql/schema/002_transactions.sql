-- +goose Up
CREATE TABLE transactions(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name TEXT NOT NULL UNIQUE,
    type TEXT NOT NULL,
    amount FLOAT NOT NULL,
    due_date INT NOT NULL,
    bank TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE transactions;
