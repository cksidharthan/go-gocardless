package requisitions_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen"
	"github.com/weportfolio/go-nordigen/consts"
	"github.com/weportfolio/go-nordigen/endpoints/agreements"
	"github.com/weportfolio/go-nordigen/endpoints/requisitions"
	"math/rand"
	"strconv"
	"testing"
)

func TestClient_Post(t *testing.T) {
	t.Parallel()

	t.Run("create new requisition", func(t *testing.T) {
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
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &requisitions.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     consts.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      consts.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.Requisitions().Post(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, consts.TestInstitutionID, requisition.InstitutionID)
	})

	t.Run("create new requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		requisitionRequestBody := &requisitions.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     consts.TestInstitutionID,
			Agreement:         "invalid",
			Reference:         "12345",
			UserLanguage:      consts.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.Requisitions().Post(context.Background(), "invalid", requisitionRequestBody)
		assert.Error(t, err)
		assert.Nil(t, requisition)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_List(t *testing.T) {
	t.Parallel()

	t.Run("list requisitions", func(t *testing.T) {
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
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &requisitions.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     consts.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      consts.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.Requisitions().Post(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, consts.TestInstitutionID, requisition.InstitutionID)

		responseRequisitions, err := client.Requisitions().List(context.Background(), token.Access, nil)
		assert.NoError(t, err)
		assert.NotNil(t, responseRequisitions)
	})

	t.Run("list requisitions with invalid token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		responseRequisitions, err := client.Requisitions().List(context.Background(), "invalid", nil)
		assert.Error(t, err)
		assert.Nil(t, responseRequisitions)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_Fetch(t *testing.T) {
	t.Parallel()

	t.Run("fetch requisition", func(t *testing.T) {
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
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &requisitions.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     consts.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      consts.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.Requisitions().Post(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, consts.TestInstitutionID, requisition.InstitutionID)

		responseRequisition, err := client.Requisitions().Fetch(context.Background(), token.Access, requisition.ID)
		assert.NoError(t, err)
		assert.NotNil(t, responseRequisition)
	})

	t.Run("fetch requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		responseRequisition, err := client.Requisitions().Fetch(context.Background(), "invalid", "invalid")
		assert.Error(t, err)
		assert.Nil(t, responseRequisition)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_Delete(t *testing.T) {
	t.Parallel()

	t.Run("delete requisition", func(t *testing.T) {
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
		assert.Equal(t, consts.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &requisitions.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     consts.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      consts.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.Requisitions().Post(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, consts.TestInstitutionID, requisition.InstitutionID)

		err = client.Requisitions().Delete(context.Background(), token.Access, requisition.ID)
		assert.NoError(t, err)
	})

	t.Run("delete requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client := nordigen.New(
			consts.GetSecrets(t),
		)
		assert.NotNil(t, client)

		err := client.Requisitions().Delete(context.Background(), "invalid", "invalid")
		assert.Error(t, err)

		checkErr := consts.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
