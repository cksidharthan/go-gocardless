package gocardless_test

import (
	"context"
	"testing"

	gocardless "github.com/cksidharthan/go-gocardless"
	"github.com/stretchr/testify/assert"
)

func TestClient_ListInstitutions(t *testing.T) {
	t.Parallel()

	t.Run("list institutions", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		institutions, err := client.ListInstitutions(context.Background(), gocardless.NetherlandsInstitution, true)
		assert.NoError(t, err)
		assert.NotNil(t, institutions)
	})

	t.Run("list institutions with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		institutions, err := client.ListInstitutions(context.Background(), gocardless.NetherlandsInstitution, true)
		assert.Error(t, err)
		assert.Nil(t, institutions)

		checkErr := gocardless.ExtractError(err)
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

		institution, err := client.FetchInstitution(context.Background(), gocardless.TestInstitutionID)
		assert.NoError(t, err)
		assert.NotNil(t, institution)
		assert.Equal(t, gocardless.TestInstitutionID, institution.ID)
	})

	t.Run("fetch institution with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		institution, err := client.FetchInstitution(context.Background(), gocardless.TestInstitutionID)
		assert.Error(t, err)
		assert.Nil(t, institution)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
