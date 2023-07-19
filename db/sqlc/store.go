package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

type TransferTxParam struct {
	FromAccountID 	int64 	`json:"from_account_id"`
	ToAccountID 	int64 	`json:"to_account_id"`
	Amount 			int64 	`json:"amount"`
}

type TransferTxResult struct {
	Transfer 		Transfer 	`json:"transfer"`
	FromAccount 	Account 	`json:"from_account"`
	ToAccount 		Account 	`json:"to_account"`
	FromEntry 		Entry 		`json:"from_entry"`
	ToEntry 		Entry 		`json:"to_entry"`
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: 		db,
		Queries: 	New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(db *Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return fmt.Errorf("tx err: %v, rErr: %v", err, rErr);
		}
		return err
	}

	return tx.Commit()
}

func (store *Store) TransferTx(ctx context.Context, args TransferTxParam) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: 	args.FromAccountID,
			ToAccountID: 	args.ToAccountID,
			Amount: 		args.Amount,
		})
		if err != nil {
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: 		args.FromAccountID,
			Amount: 		-args.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: 		args.ToAccountID,
			Amount: 		args.Amount,
		})
		if err != nil {
			return err
		}

		//TODO: deadlock impl
  
		return nil
	})
	
	return result, err
}