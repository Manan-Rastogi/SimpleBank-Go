// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0
// source: transfer.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id, 
  to_account_id, 
  amount
) VALUES (
  $1, $2, $3
)
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferByFromAccountID = `-- name: GetTransferByFromAccountID :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = $1 LIMIT 1
`

func (q *Queries) GetTransferByFromAccountID(ctx context.Context, fromAccountID int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferByFromAccountID, fromAccountID)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferByFromAndToAccountID = `-- name: GetTransferByFromAndToAccountID :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = $1 AND to_account_id = $2 LIMIT 1
`

type GetTransferByFromAndToAccountIDParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
}

func (q *Queries) GetTransferByFromAndToAccountID(ctx context.Context, arg GetTransferByFromAndToAccountIDParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferByFromAndToAccountID, arg.FromAccountID, arg.ToAccountID)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferByID = `-- name: GetTransferByID :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransferByID(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferByID, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferByToAccountID = `-- name: GetTransferByToAccountID :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE to_account_id = $1 LIMIT 1
`

func (q *Queries) GetTransferByToAccountID(ctx context.Context, toAccountID int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransferByToAccountID, toAccountID)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const listTransfersByFromAccountID = `-- name: ListTransfersByFromAccountID :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY from_account_id
Limit $1
Offset $2
`

type ListTransfersByFromAccountIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfersByFromAccountID(ctx context.Context, arg ListTransfersByFromAccountIDParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByFromAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByFromAndToAccountId = `-- name: ListTransfersByFromAndToAccountId :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY from_account_id = $1 AND to_account_id = $2
Limit $1
Offset $2
`

type ListTransfersByFromAndToAccountIdParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfersByFromAndToAccountId(ctx context.Context, arg ListTransfersByFromAndToAccountIdParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByFromAndToAccountId, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByID = `-- name: ListTransfersByID :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
Limit $1
Offset $2
`

type ListTransfersByIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfersByID(ctx context.Context, arg ListTransfersByIDParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTransfersByToAccountID = `-- name: ListTransfersByToAccountID :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY to_account_id
Limit $1
Offset $2
`

type ListTransfersByToAccountIDParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTransfersByToAccountID(ctx context.Context, arg ListTransfersByToAccountIDParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfersByToAccountID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}