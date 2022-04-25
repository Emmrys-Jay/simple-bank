package db

import (
	"context"
	"testing"

	"github.com/Emmrys-Jay/simple-bank/util"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T) Entry {
	account1 := createRandomAccount(t)

	arg := CreateEntryParams{
		AccountID: account1.ID,
		Amount:    util.RandomEntryAmount(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, account1.ID, entry.AccountID)
	require.NotEmpty(t, entry.ID)
	require.NotEmpty(t, entry.Amount)
	require.NotEmpty(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	createRandomEntry(t)
}

func TestGetEntry(t *testing.T) {
	entry1 := createRandomEntry(t)

	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
	require.Equal(t, entry1.CreatedAt, entry2.CreatedAt)
}

func TestListEntries(t *testing.T) {
	account := createRandomAccount(t)
	for i := 0; i < 10; i++ {
		arg := CreateEntryParams{
			AccountID: account.ID,
			Amount:    util.RandomEntryAmount(),
		}
		testQueries.CreateEntry(context.Background(), arg)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit:     5,
		Offset:    5,
	}
	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 5)

	for i := 0; i < 5; i++ {
		require.Equal(t, account.ID, entries[i].AccountID)
		require.NotEmpty(t, entries[i])
	}
}
