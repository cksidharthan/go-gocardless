package gocardless_test

import (
	"context"
	"testing"

	nordigen "github.com/cksidharthan/go-gocardless"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListInstitutions(t *testing.T) {
	t.Parallel()

	t.Run("list institutions", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		institutions, err := client.ListInstitutions(context.Background(), token.Access, nordigen.NetherlandsInstitution, true)
		assert.NoError(t, err)
		assert.NotNil(t, institutions)
	})

	t.Run("list institutions with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		institutions, err := client.ListInstitutions(context.Background(), "invalid", nordigen.NetherlandsInstitution, true)
		assert.Error(t, err)
		assert.Nil(t, institutions)

		checkErr := nordigen.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_FetchInstitution(t *testing.T) {
	t.Parallel()

	t.Run("fetch institution", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		institution, err := client.FetchInstitution(context.Background(), token.Access, nordigen.TestInstitutionID)
		assert.NoError(t, err)
		assert.NotNil(t, institution)
		assert.Equal(t, nordigen.TestInstitutionID, institution.ID)
	})

	t.Run("fetch institution with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		institution, err := client.FetchInstitution(context.Background(), "invalid", nordigen.TestInstitutionID)
		assert.Error(t, err)
		assert.Nil(t, institution)

		checkErr := nordigen.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
