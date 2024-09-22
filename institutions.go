package gocardless

// ListInstitutionsParams are the parameters that can be used when listing institutions
// NOTE: Except for Country, all other fields should be boolean strings
type ListInstitutionsParams struct {
	AccessScopesSupported        string `url:"access_scopes_supported,omitempty" json:"access_scopes_supported,omitempty"`
	AccountSelectionSupported    string `url:"account_selection_supported,omitempty" json:"account_selection_supported,omitempty"`
	BusinessAccountsSupported    string `url:"business_accounts_supported,omitempty" json:"business_accounts_supported,omitempty"`
	CardAccountsSupported        string `url:"card_accounts_supported,omitempty" json:"card_accounts_supported,omitempty"`
	CorporateAccountsSupported   string `url:"corporate_accounts_supported,omitempty" json:"corporate_accounts_supported,omitempty"`
	Country                      string `url:"country,omitempty" json:"country,omitempty"`
	PaymentSubmissionSupported   string `url:"payment_submission_supported,omitempty" json:"payment_submission_supported,omitempty"`
	PaymentsEnabled              string `url:"payments_enabled,omitempty" json:"payments_enabled,omitempty"`
	PendingTransactionsSupported string `url:"pending_transactions_supported,omitempty" json:"pending_transactions_supported,omitempty"`
	PrivateAccountsSupported     string `url:"private_accounts_supported,omitempty" json:"private_accounts_supported,omitempty"`
	ReadDebtorAccountSupported   string `url:"read_debtor_account_supported,omitempty" json:"read_debtor_account_supported,omitempty"`
	ReadRefundAccountSupported   string `url:"read_refund_account_supported,omitempty" json:"read_refund_account_supported,omitempty"`
	SSNVerificationSupported     string `url:"ssn_verification_supported,omitempty" json:"ssn_verification_supported,omitempty"`
}

type Institution struct {
	ID                   string            `json:"id"`
	Name                 string            `json:"name"`
	BIC                  string            `json:"bic"`
	TransactionTotalDays string            `json:"transaction_total_days"`
	Countries            []string          `json:"countries"`
	Logo                 string            `json:"logo"`
	SupportedPayments    SupportedPayments `json:"supported_payments"`
	SupportedFeatures    []string          `json:"supported_features"`
	IdentificationCodes  []string          `json:"identification_codes"`
}

type SupportedPayments struct {
	SinglePayment []string `json:"single_payment"`
}
