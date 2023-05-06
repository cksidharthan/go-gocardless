package accounts_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen"
	"github.com/weportfolio/go-nordigen/consts"
	"os"
	"testing"
)

func TestClient_GetAccount(t *testing.T) {
	t.Parallel()

	t.Run("get an account by ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		account, err := client.Accounts().GetAccount(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, account)
	})

	t.Run("get an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		account, err := client.Accounts().GetAccount(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, account)
	})
}

func TestClient_GetBalances(t *testing.T) {
	t.Parallel()

	t.Run("get balances for an account by ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		balances, err := client.Accounts().GetBalances(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, balances)
	})

	t.Run("get balances for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		balances, err := client.Accounts().GetBalances(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, balances)
	})
}

func TestClient_GetDetails(t *testing.T) {
	t.Parallel()

	t.Run("get details for an account by ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		details, err := client.Accounts().GetDetails(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, details)
	})

	t.Run("get details for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		details, err := client.Accounts().GetDetails(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, details)
	})
}

func TestClient_GetTransactions(t *testing.T) {
	t.Parallel()

	t.Run("get transactions for an account by ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		testAccountID := os.Getenv("NORDIGEN_TEST_ACCOUNT_ID")

		transactions, err := client.Accounts().GetTransactions(context.Background(), token.Access, testAccountID)
		assert.NoError(t, err)
		assert.NotNil(t, transactions)
	})

	t.Run("get transactions for an account by invalid ID", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		transactions, err := client.Accounts().GetTransactions(context.Background(), token.Access, "invalid")
		assert.Error(t, err)
		assert.Nil(t, transactions)
	})
}
