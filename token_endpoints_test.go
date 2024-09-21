package gocardless_test

import (
	"context"
	"net/http"
	"testing"

	gocardless "github.com/cksidharthan/go-gocardless"
	"github.com/stretchr/testify/assert"
)

func TestClient_NewToken(t *testing.T) {
	t.Parallel()

	t.Run("create a new client token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)
	})

	t.Run("create a new client token with invalid secret id", func(t *testing.T) {
		t.Parallel()

		invalidClient, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, invalidClient)

		token, err := invalidClient.NewToken(context.Background())
		assert.Error(t, err)
		assert.Nil(t, token)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, http.StatusUnauthorized, checkErr.StatusCode)
	})
}

func TestClient_Refresh(t *testing.T) {
	t.Parallel()

	t.Run("refresh a client token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		refreshedToken, err := client.RefreshToken(context.Background(), token.Refresh)
		assert.NoError(t, err)
		assert.NotNil(t, refreshedToken)
	})

	t.Run("refresh a client token with invalid refresh token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		refreshedToken, err := client.RefreshToken(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Nil(t, refreshedToken)
	})
}
