package nordigen_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/weportfolio/go-nordigen"
)

func TestClient_CreateRequisition(t *testing.T) {
	t.Parallel()

	t.Run("create new requisition", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := nordigen.AgreementRequestBody{
			InstitutionID:      nordigen.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, nordigen.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &nordigen.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     nordigen.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      nordigen.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, nordigen.TestInstitutionID, requisition.InstitutionID)
	})

	t.Run("create new requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		requisitionRequestBody := &nordigen.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     nordigen.TestInstitutionID,
			Agreement:         "invalid",
			Reference:         "12345",
			UserLanguage:      nordigen.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), "invalid", requisitionRequestBody)
		assert.Error(t, err)
		assert.Nil(t, requisition)

		checkErr := nordigen.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_ListRequisitions(t *testing.T) {
	t.Parallel()

	t.Run("list requisitions", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := nordigen.AgreementRequestBody{
			InstitutionID:      nordigen.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, nordigen.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &nordigen.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     nordigen.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      nordigen.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, nordigen.TestInstitutionID, requisition.InstitutionID)

		responseRequisitions, err := client.ListRequisitions(context.Background(), token.Access, nil)
		assert.NoError(t, err)
		assert.NotNil(t, responseRequisitions)
	})

	t.Run("list requisitions with invalid token", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		responseRequisitions, err := client.ListRequisitions(context.Background(), "invalid", nil)
		assert.Error(t, err)
		assert.Nil(t, responseRequisitions)

		checkErr := nordigen.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_FetchRequisition(t *testing.T) {
	t.Parallel()

	t.Run("fetch requisition", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := nordigen.AgreementRequestBody{
			InstitutionID:      nordigen.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, nordigen.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &nordigen.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     nordigen.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      nordigen.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, nordigen.TestInstitutionID, requisition.InstitutionID)

		responseRequisition, err := client.FetchRequisition(context.Background(), token.Access, requisition.ID)
		assert.NoError(t, err)
		assert.NotNil(t, responseRequisition)
	})

	t.Run("fetch requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		responseRequisition, err := client.FetchRequisition(context.Background(), "invalid", "invalid")
		assert.Error(t, err)
		assert.Nil(t, responseRequisition)

		checkErr := nordigen.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_DeleteRequisition(t *testing.T) {
	t.Parallel()

	t.Run("delete requisition", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		token, err := client.NewToken(context.Background())
		assert.NoError(t, err)
		assert.NotNil(t, token)

		agreementRequestBody := nordigen.AgreementRequestBody{
			InstitutionID:      nordigen.TestInstitutionID,
			MaxHistoricalDays:  "180",
			AccessValidForDays: "2",
			AccessScope:        []string{"balances", "details", "transactions"},
		}

		agreement, err := client.CreateAgreement(context.Background(), token.Access, agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, nordigen.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &nordigen.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     nordigen.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      nordigen.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), token.Access, requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, nordigen.TestInstitutionID, requisition.InstitutionID)

		err = client.DeleteRequisition(context.Background(), token.Access, requisition.ID)
		assert.NoError(t, err)
	})

	t.Run("delete requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client := getTestClient(t)
		assert.NotNil(t, client)

		err := client.DeleteRequisition(context.Background(), "invalid", "invalid")
		assert.Error(t, err)

		checkErr := nordigen.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
