package gocardless

import (
	"context"
	"fmt"
	"net/url"
)

// ListInstitutions returns a list of institutions
func (c Client) ListInstitutions(ctx context.Context, queryParams ListInstitutionsParams) ([]Institution, error) {
	endpointURL, err := url.Parse(InstitutionsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %v", err)
	}

	query := bindListInstitutionsParams(queryParams)

	endpointURL.RawQuery = query.Encode()

	var institutions []Institution

	err = c.HTTP.Get(ctx, endpointURL.String(), RequestHeadersWithAuth(c.Token.Access), &institutions)
	if err != nil {
		return nil, err
	}

	return institutions, nil
}

// FetchInstitution returns an institution
func (c Client) FetchInstitution(ctx context.Context, id string) (*Institution, error) {
	var institution Institution
	err := c.HTTP.Get(ctx, InstitutionsPath+id, RequestHeadersWithAuth(c.Token.Access), &institution)
	if err != nil {
		return nil, err
	}

	return &institution, nil
}

func bindListInstitutionsParams(queryParams ListInstitutionsParams) url.Values {
	query := make(url.Values)

	if queryParams.Country != "" {
		query.Add("country", queryParams.Country)
	}
	if queryParams.AccessScopesSupported != "" {
		query.Add("access_scopes_supported", queryParams.AccessScopesSupported)
	}
	if queryParams.AccountSelectionSupported != "" {
		query.Add("account_selection_supported", queryParams.AccountSelectionSupported)
	}
	if queryParams.BusinessAccountsSupported != "" {
		query.Add("business_accounts_supported", queryParams.BusinessAccountsSupported)
	}
	if queryParams.CardAccountsSupported != "" {
		query.Add("card_accounts_supported", queryParams.CardAccountsSupported)
	}
	if queryParams.CorporateAccountsSupported != "" {
		query.Add("corporate_accounts_supported", queryParams.CorporateAccountsSupported)
	}
	if queryParams.PaymentSubmissionSupported != "" {
		query.Add("payment_submission_supported", queryParams.PaymentSubmissionSupported)
	}
	if queryParams.PaymentsEnabled != "" {
		query.Add("payments_enabled", queryParams.PaymentsEnabled)
	}
	if queryParams.PendingTransactionsSupported != "" {
		query.Add("pending_transactions_supported", queryParams.PendingTransactionsSupported)
	}
	if queryParams.PrivateAccountsSupported != "" {
		query.Add("private_accounts_supported", queryParams.PrivateAccountsSupported)
	}
	if queryParams.ReadDebtorAccountSupported != "" {
		query.Add("read_debtor_account_supported", queryParams.ReadDebtorAccountSupported)
	}
	if queryParams.ReadRefundAccountSupported != "" {
		query.Add("read_refund_account_supported", queryParams.ReadRefundAccountSupported)
	}
	if queryParams.SSNVerificationSupported != "" {
		query.Add("ssn_verification_supported", queryParams.SSNVerificationSupported)
	}
	return query
}
