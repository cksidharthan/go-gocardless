package token_test

import (
	"fmt"
	"github.com/weportfolio/go-nordigen/consts"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen"
)

func TestClient_New(t *testing.T) {
	t.Parallel()

	t.Run("create a new client token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			getSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New()
		assert.NoError(t, err)
		assert.NotNil(t, token)
	})

	t.Run("create a new client token with invalid secret id", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			"invalid",
			"invalid",
		)
		assert.NotNil(t, client)

		token, err := client.Token().New()
		assert.Error(t, err)
		assert.Nil(t, token)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, http.StatusUnauthorized, checkErr.StatusCode)
	})
}

func TestClient_Refresh(t *testing.T) {
	t.Parallel()

	t.Run("refresh a client token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			getSecrets(t),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New()
		assert.NoError(t, err)
		assert.NotNil(t, token)

		refreshedToken, err := client.Token().Refresh(token.Refresh)
		assert.NoError(t, err)
		assert.NotNil(t, refreshedToken)
	})

	t.Run("refresh a client token with invalid refresh token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			getSecrets(t),
		)
		assert.NotNil(t, client)

		refreshedToken, err := client.Token().Refresh("invalid")
		assert.Error(t, err)
		assert.Nil(t, refreshedToken)
	})
}

func getSecrets(t *testing.T) (string, string) {
	t.Helper()

	secretID := os.Getenv("NORDIGEN_SECRET_ID")
	secretKey := os.Getenv("NORDIGEN_SECRET_KEY")
	if secretID == "" || secretKey == "" {
		fmt.Println("NORDIGEN_SECRET_ID or NORDIGEN_SECRET_KEY is not set")
	}

	return secretID, secretKey
}
