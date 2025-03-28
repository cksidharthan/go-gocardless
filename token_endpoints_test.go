package gocardless_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient_NewToken(t *testing.T) {
	t.Parallel()

	t.Run("create a new client token", func(t *testing.T) {
		t.Parallel()

		token, err := testClient.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)
	})

	t.Run("create a new client token with invalid secret id", func(t *testing.T) {
		t.Parallel()

		invalidClient, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, invalidClient)
	})
}

func TestClient_Refresh(t *testing.T) {
	t.Parallel()

	t.Run("refresh a client token", func(t *testing.T) {
		t.Parallel()

		token, err := testClient.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		refreshedToken, err := testClient.RefreshToken(context.Background(), token.Refresh)
		assert.NoError(t, err)
		assert.NotNil(t, refreshedToken)
	})

	t.Run("refresh a client token with invalid refresh token", func(t *testing.T) {
		t.Parallel()

		refreshedToken, err := testClient.RefreshToken(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Nil(t, refreshedToken)
	})
}
