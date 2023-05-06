package accounts

import "time"

type Account struct {
	ID            string `json:"id"`
	Created       string `json:"created"`
	LastAccessed  string `json:"last_accessed"`
	IBAN          string `json:"iban"`
	InstitutionID string `json:"institution_id"`
	Status        string `json:"status"`
	OwnerName     string `json:"owner_name"`
}

type Balance struct {
	BalanceAmount Amount `json:"balanceAmount"`
	BalanceType   string `json:"balanceType"`
	ReferenceDate string `json:"referenceDate"`
}

type Amount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type Balances struct {
	Balances []Balance `json:"balances"`
}

type AccountDetails struct {
	ResourceID      string `json:"resourceId"`
	IBAN            string `json:"iban"`
	Currency        string `json:"currency"`
	OwnerName       string `json:"ownerName"`
	Name            string `json:"name"`
	Product         string `json:"product"`
	CashAccountType string `json:"cashAccountType"`
}

type Details struct {
	Account AccountDetails `json:"account"`
}

type TransactionParams struct {
	DateFrom string `url:"date_from,omitempty" json:"date_from,omitempty"`
	DateTo   string `url:"date_to,omitempty" json:"date_to,omitempty"`
}

type Transaction struct {
	TransactionID                          string    `json:"transactionId"`
	BookingDate                            string    `json:"bookingDate"`
	ValueDate                              string    `json:"valueDate"`
	BookingDateTime                        time.Time `json:"bookingDateTime"`
	ValueDateTime                          time.Time `json:"valueDateTime"`
	TransactionAmount                      Amount    `json:"transactionAmount"`
	CreditorName                           string    `json:"creditorName"`
	CreditorAccount                        Account   `json:"creditorAccount"`
	DebtorName                             string    `json:"debtorName"`
	DebtorAccount                          Account   `json:"debtorAccount"`
	BankTransactionCode                    string    `json:"bankTransactionCode"`
	RemittanceInformationUnstructured      string    `json:"remittanceInformationUnstructured"`
	RemittanceInformationUnstructuredArray []string  `json:"remittanceInformationUnstructuredArray"`
	ProprietaryBankTransactionCode         string    `json:"proprietaryBankTransactionCode"`
	InternalTransactionID                  string    `json:"internalTransactionId"`
}

type Transactions struct {
	Transactions TransactionList `json:"transactions"`
}

type TransactionList struct {
	Booked  []Transaction `json:"booked"`
	Pending []Transaction `json:"pending"`
}
