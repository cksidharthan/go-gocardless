package nordigen

type Institution struct {
	ID                   string            `json:"id"`
	Name                 string            `json:"name"`
	BIC                  string            `json:"bic"`
	TransactionTotalDays string            `json:"transaction_total_days"`
	Countries            []string          `json:"countries"`
	Logo                 string            `json:"logo"`
	SupportedPayments    SupportedPayments `json:"supported_payments"`
	SupportedFeatures    []string          `json:"supported_features"`
}

type SupportedPayments struct {
	SinglePayment []string `json:"single_payment"`
}
