package gocardless_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"

	"github.com/cksidharthan/go-gocardless"
	"github.com/stretchr/testify/assert"
)

func TestClient_CreateRequisition(t *testing.T) {
	t.Parallel()

	t.Run("create new requisition", func(t *testing.T) {
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

		agreement, err := client.CreateAgreement(context.Background(), agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &gocardless.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     gocardless.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      gocardless.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, gocardless.TestInstitutionID, requisition.InstitutionID)
	})

	t.Run("create new requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		requisitionRequestBody := &gocardless.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     gocardless.TestInstitutionID,
			Agreement:         "invalid",
			Reference:         "12345",
			UserLanguage:      gocardless.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), requisitionRequestBody)
		assert.Error(t, err)
		assert.Nil(t, requisition)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_ListRequisitions(t *testing.T) {
	t.Parallel()

	t.Run("list requisitions", func(t *testing.T) {
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

		agreement, err := client.CreateAgreement(context.Background(), agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &gocardless.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     gocardless.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      gocardless.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, gocardless.TestInstitutionID, requisition.InstitutionID)

		responseRequisitions, err := client.ListRequisitions(context.Background(), nil)
		assert.NoError(t, err)
		assert.NotNil(t, responseRequisitions)
	})

	t.Run("list requisitions with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		responseRequisitions, err := client.ListRequisitions(context.Background(), nil)
		assert.Error(t, err)
		assert.Nil(t, responseRequisitions)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_FetchRequisition(t *testing.T) {
	t.Parallel()

	t.Run("fetch requisition", func(t *testing.T) {
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

		agreement, err := client.CreateAgreement(context.Background(), agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &gocardless.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     gocardless.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      gocardless.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, gocardless.TestInstitutionID, requisition.InstitutionID)

		responseRequisition, err := client.FetchRequisition(context.Background(), requisition.ID)
		assert.NoError(t, err)
		assert.NotNil(t, responseRequisition)
	})

	t.Run("fetch requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		responseRequisition, err := client.FetchRequisition(context.Background(), "invalid")
		assert.Error(t, err)
		assert.Nil(t, responseRequisition)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}

func TestClient_DeleteRequisition(t *testing.T) {
	t.Parallel()

	t.Run("delete requisition", func(t *testing.T) {
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

		agreement, err := client.CreateAgreement(context.Background(), agreementRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, agreement)
		assert.Equal(t, gocardless.TestInstitutionID, agreement.InstitutionID)

		requisitionRequestBody := &gocardless.RequisitionRequestBody{
			Redirect:          "https://example.com",
			InstitutionID:     gocardless.TestInstitutionID,
			Agreement:         agreement.ID,
			Reference:         strconv.Itoa(rand.Intn(1000000000000000000)),
			UserLanguage:      gocardless.LangEN,
			AccountSelection:  false,
			RedirectImmediate: false,
		}

		requisition, err := client.CreateRequisition(context.Background(), requisitionRequestBody)
		assert.NoError(t, err)
		assert.NotNil(t, requisition)
		assert.Equal(t, gocardless.TestInstitutionID, requisition.InstitutionID)

		err = client.DeleteRequisition(context.Background(), requisition.ID)
		assert.NoError(t, err)
	})

	t.Run("delete requisition with invalid token", func(t *testing.T) {
		t.Parallel()

		client, err := getInvalidTestClient(t)
		assert.NoError(t, err)
		assert.NotNil(t, client)

		err = client.DeleteRequisition(context.Background(), "invalid")
		assert.Error(t, err)

		checkErr := gocardless.ExtractError(err)
		assert.Equal(t, 401, checkErr.StatusCode)
	})
}
