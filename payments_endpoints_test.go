package gocardless_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_ListPayments(t *testing.T) {
	t.Parallel()

	// Payments are not enabled for the account used in the tests
	t.Run("list payments", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		payments, err := client.ListPayments(context.Background(), 10, 0)
		assert.Error(t, err)
		assert.Nil(t, payments)
	})
}
