package agreements_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen"
	"github.com/weportfolio/go-nordigen/consts"
	"github.com/weportfolio/go-nordigen/endpoints/agreements"
	"testing"
)

func TestClient_Post(t *testing.T) {
	t.Parallel()

	t.Run("create a new agreement", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
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
	})

	t.Run("create a new agreement with invalid token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
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
