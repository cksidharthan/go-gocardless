package token_test

import (
	"fmt"
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
			getSecrets(),
		)
		assert.NotNil(t, client)

		token, err := client.Token().New()
		assert.NoError(t, err)
		assert.NotNil(t, token)
	})
}

func getSecrets() (string, string) {
	secretID := os.Getenv("NORDIGEN_SECRET_ID")
	secretKey := os.Getenv("NORDIGEN_SECRET_KEY")
	if secretID == "" || secretKey == "" {
		fmt.Println("NORDIGEN_SECRET_ID or NORDIGEN_SECRET_KEY is not set")
	}

	return secretID, secretKey
}
