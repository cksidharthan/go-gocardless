package gocardless_test

import (
	"context"
	"testing"

	gocardless "github.com/cksidharthan/go-gocardless"
	"github.com/stretchr/testify/assert"
)

func TestClient_CreateAgreement(t *testing.T) {
	t.Parallel()

	t.Run("create a new agreement", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := gocardless.AgreementRequestBody{
			InstitutionID:      gocardless.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)
	})

	t.Run("create a new agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		agreementRequestBody := gocardless.AgreementRequestBody{
			InstitutionID:      gocardless.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), "invalid", agreementRequestBody)
		assert.Error(t, err)
		assert.Nil(t, agreement)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_FetchAgreement(t *testing.T) {
	t.Parallel()

	t.Run("fetch an agreement", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := gocardless.AgreementRequestBody{
			InstitutionID:      gocardless.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)

		fetchedAgreement, err := client.FetchAgreement(context.Background(), token.Access, agreement.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedAgreement)
		assert.Equal(t, agreement.ID, fetchedAgreement.ID)
	})

	t.Run("fetch an agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		agreementRequestBody := gocardless.AgreementRequestBody{
			InstitutionID:      gocardless.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), "invalid", agreementRequestBody)
		assert.Error(t, err)
		assert.Nil(t, agreement)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_ListAgreements(t *testing.T) {
	t.Parallel()

	t.Run("list agreements", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		responseAgreements, err := client.ListAgreements(context.Background(), token.Access, nil)
		assert.NoError(t, err)
		assert.NotNil(t, responseAgreements)
	})

	t.Run("list agreements with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		responseAgreements, err := client.ListAgreements(context.Background(), "invalid", nil)
		assert.Error(t, err)
		assert.Nil(t, responseAgreements)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_DeleteAgreement(t *testing.T) {
	t.Parallel()

	t.Run("delete an agreement", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := gocardless.AgreementRequestBody{
			InstitutionID:      gocardless.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)

		err = client.DeleteAgreement(context.Background(), token.Access, agreement.ID)
		assert.NoError(t, err)
	})

	t.Run("delete an agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		err = client.DeleteAgreement(context.Background(), "invalid", "invalid")
		assert.Error(t, err)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_UpdateAgreement(t *testing.T) {
	t.Parallel()

	t.Run("update an agreement", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := gocardless.AgreementRequestBody{
			InstitutionID:      gocardless.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)

		updateRequestBody := gocardless.UpdateRequestBody{
			UserAgent: "test",
			IPAddress: "0.0.0.0",
		}

		updatedAgreement, err := client.UpdateAgreement(context.Background(), token.Access, agreement.ID, updateRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, updatedAgreement)
		assert.Equal(t, gocardless.TestInstitutionID, updatedAgreement.InstitutionID)
		assert.Equal(t, agreement.ID, updatedAgreement.ID)
	})

	t.Run("update an agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		updateRequestBody := gocardless.UpdateRequestBody{
			UserAgent: "test",
			IPAddress: "",
		}

		updatedAgreement, err := client.UpdateAgreement(context.Background(), "invalid", "invalid", updateRequestBody)
		assert.Error(t, err)
		assert.Nil(t, updatedAgreement)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
