package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTransferAccounts(t *testing.T) []Account {
	account1 := CreateRandomAccount(t)
	account2 := CreateRandomAccount(t)
	
	accounts := []Account{account1, account2}
	return accounts
}

func createRandomTransfer(t *testing.T, accounts []Account) Transfer {
	params := CreateTransferParams{
		FromAccountID: accounts[0].ID,
		ToAccountID: accounts[1].ID,
		Amount: int64(10),
	}
	result, err := testQueries.CreateTransfer(context.Background(), params)
	require.NoError(t, err)
	require.NotEmpty(t, result)	
	require.Equal(t, accounts[0].ID, result.FromAccountID)
	require.Equal(t, accounts[1].ID, result.ToAccountID)
	require.Equal(t, params.Amount, result.Amount)
	require.NotZero(t, result.ID)
	require.NotZero(t, result.CreatedAt)

	return result
}

func TestCreateTransfer(t *testing.T) {
	createRandomTransfer(t, createTransferAccounts(t))
}

func TestGetTransfer(t *testing.T) {
	transfer := createRandomTransfer(t, createTransferAccounts(t))

	result, err := testQueries.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)
	require.NotEmpty(t, result)
	require.Equal(t, result.ID, transfer.ID)
	require.Equal(t, result.FromAccountID, transfer.FromAccountID)
	require.Equal(t, result.ToAccountID, transfer.ToAccountID)
	require.Equal(t, result.Amount, transfer.Amount)
}

func TestListTransfer(t * testing.T) {
	accounts := createTransferAccounts(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, accounts)
	}

	args := ListTransfersParams{
		FromAccountID: accounts[0].ID,
		ToAccountID: accounts[1].ID,
		Offset: 5,
		Limit: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), args)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
