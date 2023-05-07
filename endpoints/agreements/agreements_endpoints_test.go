package agreements_test

import (
	"context"
	"testing"

	"github.com/weportfolio/go-nordigen/endpoints/tests"

	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen/consts"
	"github.com/weportfolio/go-nordigen/endpoints/agreements"
)

func TestClient_Post(t *testing.T) {
	t.Parallel()

	t.Run("create a new agreement", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := agreements.AgreementRequestBody{
			InstitutionID:      consts.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.Agreements().Post(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)
	})

	t.Run("create a new agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		agreementRequestBody := agreements.AgreementRequestBody{
			InstitutionID:      consts.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.Agreements().Post(context.Background(), "invalid", agreementRequestBody)
		assert.Error(t, err)
		assert.Nil(t, agreement)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_Fetch(t *testing.T) {
	t.Parallel()

	t.Run("fetch an agreement", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := agreements.AgreementRequestBody{
			InstitutionID:      consts.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.Agreements().Post(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)

		fetchedAgreement, err := client.Agreements().Fetch(context.Background(), token.Access, agreement.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedAgreement)
		assert.Equal(t, agreement.ID, fetchedAgreement.ID)
	})

	t.Run("fetch an agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		agreementRequestBody := agreements.AgreementRequestBody{
			InstitutionID:      consts.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.Agreements().Post(context.Background(), "invalid", agreementRequestBody)
		assert.Error(t, err)
		assert.Nil(t, agreement)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_List(t *testing.T) {
	t.Parallel()

	t.Run("list agreements", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		responseAgreements, err := client.Agreements().List(context.Background(), token.Access, nil)
		assert.NoError(t, err)
		assert.NotNil(t, responseAgreements)
	})

	t.Run("list agreements with invalid token", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		responseAgreements, err := client.Agreements().List(context.Background(), "invalid", nil)
		assert.Error(t, err)
		assert.Nil(t, responseAgreements)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	t.Run("delete an agreement", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := agreements.AgreementRequestBody{
			InstitutionID:      consts.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.Agreements().Post(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)

		err = client.Agreements().Delete(context.Background(), token.Access, agreement.ID)
		assert.NoError(t, err)
	})

	t.Run("delete an agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		err := client.Agreements().Delete(context.Background(), "invalid", "invalid")
		assert.Error(t, err)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_Update(t *testing.T) {
	t.Parallel()

	t.Run("update an agreement", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		token, err := client.Token().New(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := agreements.AgreementRequestBody{
			InstitutionID:      consts.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.Agreements().Post(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)

		updateRequestBody := agreements.UpdateRequestBody{
			UserAgent: "test",
			IPAddress: "0.0.0.0",
		}

		updatedAgreement, err := client.Agreements().Update(context.Background(), token.Access, agreement.ID, updateRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, updatedAgreement)
		assert.Equal(t, consts.TestInstitutionID, updatedAgreement.InstitutionID)
		assert.Equal(t, agreement.ID, updatedAgreement.ID)
	})

	t.Run("update an agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client := tests.GetTestClient(t)
		assert.NotNil(t, client)

		updateRequestBody := agreements.UpdateRequestBody{
			UserAgent: "test",
			IPAddress: "",
		}

		updatedAgreement, err := client.Agreements().Update(context.Background(), "invalid", "invalid", updateRequestBody)
		assert.Error(t, err)
		assert.Nil(t, updatedAgreement)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
