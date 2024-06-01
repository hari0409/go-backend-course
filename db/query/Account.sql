-- name: CreateAccount :one
INSERT INTO "Account" ( -- Provide the object name
    owner, -- Provide all the required fields
    balance,
    currency
) VALUES ( -- The values are used for specifing that it will be sent
    $1, $2, $3
) RETURNING *; -- The return statement with "*" to indicate to return all fields

-- name: GetAccount :one
SELECT * FROM "Account"
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM "Account"
WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM "Account" 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE "Account" 
SET balance = $2
WHERE id = $1
RETURNING *;

-- name: AddAccountBalance :one
UPDATE "Account" 
SET balance = balance+ sqlc.arg(amount)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM "Account"
WHERE id = $1; 
