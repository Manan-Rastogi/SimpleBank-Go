-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, 
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
)
RETURNING *; 

-- name: GetTransferByID :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransferByFromAccountID :one
SELECT * FROM transfers
WHERE from_account_id = $1 LIMIT 1;

-- name: GetTransferByToAccountID :one
SELECT * FROM transfers
WHERE to_account_id = $1 LIMIT 1;

-- name: GetTransferByFromAndToAccountID :one
SELECT * FROM transfers
WHERE from_account_id = $1 AND to_account_id = $2 LIMIT 1;


-- name: ListTransfersByID :many
SELECT * FROM transfers
ORDER BY id
Limit $1
Offset $2;

-- name: ListTransfersByFromAccountID :many
SELECT * FROM transfers
ORDER BY from_account_id
Limit $1
Offset $2;

-- name: ListTransfersByToAccountID :many
SELECT * FROM transfers
ORDER BY to_account_id
Limit $1
Offset $2;

-- name: ListTransfersByFromAndToAccountId :many
SELECT * FROM transfers
ORDER BY from_account_id = $1 AND to_account_id = $2
Limit $1
Offset $2;


-- -- name: UpdateTransfer :one
-- UPDATE transfers 
-- SET balance = $2
-- WHERE id = $1
-- RETURNING *;

-- -- name: DeleteTransfer :exec
-- DELETE FROM transfers
-- WHERE id = $1;
