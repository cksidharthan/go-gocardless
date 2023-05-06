package payments_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen"
	"github.com/weportfolio/go-nordigen/consts"
	"testing"
)

func TestClient_List(t *testing.T) {
	t.Parallel()

	// Payments are not enabled for the account used in the tests
	t.Run("list payments", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		payments, err := client.Payments().List(context.Background(), token.Access, 10, 0)
		assert.Error(t, err)
		assert.Nil(t, payments)
	})
}
