package db

import (
	"context"
	"testing"
	"time"

	"github.com/IkehAkinyemi/mono-finance/utils"
	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, accountID int64) Entry {
	arg := CreateEntryParams {
		AccountID: accountID,
		Amount: utils.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	require.Equal(t, arg.AccountID, entry.AccountID)
	require.Equal(t, arg.Amount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	account := createRandomAccount(t)
	createRandomEntry(t, account.ID)
}

func TestGetEntry(t *testing.T) {
	account := createRandomAccount(t)
	entry1 := createRandomEntry(t, account.ID)
	
	entry2, err := testQueries.GetEntry(context.Background(), entry1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entry2)

	require.Equal(t, entry1.ID, entry2.ID)
	require.WithinDuration(t, entry1.CreatedAt, entry2.CreatedAt, time.Second)
	require.Equal(t, entry1.AccountID, entry2.AccountID)
	require.Equal(t, entry1.Amount, entry2.Amount)
}

func TestListEntry(t *testing.T) {
	account := createRandomAccount(t)

	for i := 0; i < 12; i++ {
		createRandomEntry(t, account.ID)
	}

	arg := ListEntriesParams{
		AccountID: account.ID,
		Limit: 7,
		Offset: 5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, entries, 7)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
		require.Equal(t, arg.AccountID, entry.AccountID)
	}
}
