package db

import (
	"context"
	"database/sql"
	"fmt"
)

// Store provides all functions to execute db queries and transactions
type Store struct {
	*Queries
	db *sql.DB
}

// Creates a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// executes a function within a database transaction
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	// BEGIN TRANSACTION
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// QUERIES FOR TRANSACTION
	q := New(tx)
	// EXECUTE QUERIES
	err = fn(q)
	if err != nil {
		// ROLLBACK IF THERE IS ERROR
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx Error %v, rb Error %v", err, rbErr)
		}
		return err
	}

	// COMMIT IF THERE IS NO ERROR i.e SUCCESSFUL TRANSACTION
	return tx.Commit()
}

// Contains input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountId   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// Result of transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

var txKey = struct{}{}

// Performs a money transfer from 1 account to another.
// It creates a transfer record, add account entries and update account's balance within a single db transaction.
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		txName := ctx.Value(txKey)
		fmt.Println(txName, "create transfer")
		// 1. CREATE A TRANSFER RECORD.
		result.Transfer, err =  q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountId,
			ToAccountID: arg.ToAccountId,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 1")
		// 2. ADD ACCOUNT ENTRIES
		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountId,
			Amount: -arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "create entry 2")
		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountId,
			Amount: arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "get account 1")
		//  3. Update Account Balance
		fromAccountId, err := q.GetAccountForUpdate(ctx, arg.FromAccountId)
		if err != nil {
			return err
		}

		fmt.Println(txName, "get account 2")
		toAccountId, err := q.GetAccountForUpdate(ctx, arg.ToAccountId)
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 1")
		result.FromAccount, err =  q.UpdateAccount(ctx, UpdateAccountParams{
			ID: arg.FromAccountId,
			Balance: fromAccountId.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}

		fmt.Println(txName, "update account 2" )
		result.ToAccount, err =  q.UpdateAccount(ctx, UpdateAccountParams{
			ID: arg.ToAccountId,
			Balance: toAccountId.Balance + arg.Amount,
		})
		if err != nil {
			return err
		}



		
		return nil
	})

	return result, err
}
