package gocardless_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_GetAccount(t *testing.T) {
	t.Parallel()

	t.Run("get an account by ID", func(t *testing.T) {
		t.Parallel()

		account, err := testClient.GetAccount(context.Background(), testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})

	t.Run("get an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		account, err := testClient.GetAccount(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Nil(t, account)
	})
}

func TestClient_GetAccountBalances(t *testing.T) {
	t.Parallel()

	t.Run("get balances for an account by ID", func(t *testing.T) {
		t.Parallel()

		balances, err := testClient.GetAccountBalances(context.Background(), testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, balances)
	})

	t.Run("get balances for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		balances, err := testClient.GetAccountBalances(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Nil(t, balances)
	})
}

func TestClient_GetAccountDetails(t *testing.T) {
	t.Parallel()

	t.Run("get details for an account by ID", func(t *testing.T) {
		t.Parallel()

		details, err := testClient.GetAccountDetails(context.Background(), testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, details)
	})

	t.Run("get details for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		details, err := testClient.GetAccountDetails(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Nil(t, details)
	})
}

func TestClient_GetAccountTransactions(t *testing.T) {
	t.Parallel()

	t.Run("get transactions for an account by ID", func(t *testing.T) {
		t.Parallel()

		transactions, err := testClient.GetAccountTransactions(context.Background(), testAccountID, nil, nil)
		assert.NoError(t, err)
		assert.NotNil(t, transactions)
	})

	t.Run("get transactions for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		transactions, err := testClient.GetAccountTransactions(context.Background(), "invalid", nil, nil)
		assert.Error(t, err)
		assert.Nil(t, transactions)
	})
}
