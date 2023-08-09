package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) (Entry) {
	account := CreateRandomAccount(t)

	args := CreateEntryParams{
		AccountID: account.ID,
		Amount: 50,
	}
	entry, err := testQueries.CreateEntry(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, entry.AccountID, account.ID)
	require.Equal(t, args.AccountID, entry.AccountID)
	require.Equal(t, args.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestDeleteEntry(t *testing.T) {
	entry := createRandomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entry.ID)

	require.NoError(t, err)
	entry1, err := testQueries.GetEntry(context.Background(), entry.ID)

	require.Error(t, err)
	require.Empty(t, entry1)
	require.EqualError(t, err, sql.ErrNoRows.Error())
}

func TestUpdateEntry(t *testing.T) {
	entry := createRandomEntry(t)

	args := UpdateEntryParams{
		ID: entry.ID,
		Amount: 83,
	}
	entryNew, err := testQueries.UpdateEntry(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, entryNew)
	require.Equal(t, args.Amount, entryNew.Amount)
	require.Equal(t, args.ID, entryNew.ID)
	require.Equal(t, entry.ID, args.ID)
	require.WithinDuration(t, entry.CreatedAt, entryNew.CreatedAt, time.Second)
}

// func TestListEntry(t *testing.T) {
// 	for i := 0; i < 10; i++ {
// 		createRandomEntry(t)
// 	}

// 	args := ListEntriesParams{
// 		Offset: 5,
// 		Limit: 5,
// 	}

// 	entries, err := testQueries.ListEntries(context.Background(), args)

// 	require.NoError(t, err)
// 	require.Len(t, entries, 5)

// 	for _, entry := range entries {
// 		require.NotEmpty(t, entry)
// 	}
// }