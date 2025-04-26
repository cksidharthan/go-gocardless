package gocardless

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	ID            uuid.UUID `json:"id"`
	Created       time.Time `json:"created"`
	LastAccessed  time.Time `json:"last_accessed"`
	IBAN          string    `json:"iban"`
	BBAN          string    `json:"bban"`
	InstitutionID string    `json:"institution_id"`
	Status        string    `json:"status"`
	OwnerName     string    `json:"owner_name"`
}

type Balance struct {
	BalanceAmount            Amount `json:"balanceAmount"`
	BalanceType              string `json:"balanceType"`
	ReferenceDate            string `json:"referenceDate"`
	CreditLimitIncluded      bool   `json:"creditLimitIncluded"`
	LastCommittedTransaction string `json:"lastCommittedTransaction"`
}

type Amount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

type AccountBalance struct {
	Balances []Balance `json:"balances"`
}

// AccountDetail is a struct that contains the details of an account
// Some fields might be empty, depending on the account type
type AccountDetail struct {
	ResourceID               string                 `json:"resourceId"`
	IBAN                     string                 `json:"iban"`
	BBAN                     string                 `json:"bban"`
	Scan                     string                 `json:"scan"`
	Msisdn                   string                 `json:"msisdn"`
	Currency                 string                 `json:"currency"`
	OwnerName                string                 `json:"ownerName"`
	Name                     string                 `json:"name"`
	DisplayName              string                 `json:"displayName"`
	Product                  string                 `json:"product"`
	CashAccountType          string                 `json:"cashAccountType"`
	Status                   string                 `json:"status"`
	BIC                      string                 `json:"bic"`
	LinkedAccounts           string                 `json:"linkedAccounts"`
	MaskedPAN                string                 `json:"maskedPan"`
	Usage                    string                 `json:"usage"`
	Details                  string                 `json:"details"`
	OwnerAddressUnstructured string                 `json:"ownerAddressUnstructured"`
	OwnerAddressStructured   OwnerAddressStructured `json:"ownerAddressStructured"`
}

type OwnerAddressStructured struct {
	StreetName     string `json:"streetName"`
	BuildingNumber string `json:"buildingNumber"`
	Postcode       string `json:"postCode"`
	Country        string `json:"country"`
	TownName       string `json:"townName"`
}

type Details struct {
	Account AccountDetail `json:"account"`
}

type TransactionParams struct {
	DateFrom time.Time `url:"date_from,omitempty" json:"date_from,omitempty"`
	DateTo   time.Time `url:"date_to,omitempty" json:"date_to,omitempty"`
	ID       string    `url:"id,omitempty" json:"id,omitempty"`
}

type Transaction struct {
	TransactionID                          string                `json:"transactionId"`
	EntryReference                         string                `json:"entryReference"`
	EndToEndID                             string                `json:"endToEndId"`
	MandateID                              string                `json:"mandateId"`
	CheckID                                string                `json:"checkId"`
	CreditorID                             string                `json:"creditorId"`
	BookingDate                            string                `json:"bookingDate"`
	ValueDate                              string                `json:"valueDate"`
	BookingDateTime                        TimeWithTimeZoneInfoZ `json:"bookingDateTime"`
	ValueDateTime                          TimeWithTimeZoneInfoZ `json:"valueDateTime"`
	TransactionAmount                      Amount                `json:"transactionAmount"`
	CurrencyExchange                       []CurrencyExchange    `json:"currencyExchange"`
	CreditorName                           string                `json:"creditorName"`
	CreditorAccount                        Account               `json:"creditorAccount"`
	UltimateCollector                      string                `json:"ultimateCreditor"`
	RemittanceInformationUnstructured      string                `json:"remittanceInformationUnstructured"`
	RemittanceInformationUnstructuredArray []string              `json:"remittanceInformationUnstructuredArray"`
	RemittanceInformationStructured        string                `json:"remittanceInformationStructured"`
	RemittanceInformationStructuredArray   []string              `json:"remittanceInformationStructuredArray"`
	AdditionalInformation                  string                `json:"additionalInformation"`
	PurposeCode                            string                `json:"purposeCode"`
	BankTransactionCode                    string                `json:"bankTransactionCode"`
	ProprietaryBankTransactionCode         string                `json:"proprietaryBankTransactionCode"`
	InternalTransactionID                  string                `json:"internalTransactionId"`
	BalanceAfterTransaction                Amount                `json:"balanceAfterTransaction"`
	// The below fields are probably deprecated.
	DebtorName                      string  `json:"debtorName"`
	DebtorAccount                   Account `json:"debtorAccount"`
	AdditionalInformationStructured string  `json:"additionalInformationStructured"`
	DebtorAgent                     string  `json:"debtorAgent"`
	MerchantCategoryCode            string  `json:"merchantCategoryCode"`
	UltimateDebtor                  string  `json:"ultimateDebtor"`
}

type CurrencyExchange struct {
	SourceCurrency         string `json:"sourceCurrency"`
	ExchangeRate           string `json:"exchangeRate"`
	UnitCurrency           string `json:"unitCurrency"`
	TargetCurrency         string `json:"targetCurrency"`
	QuotationDate          string `json:"quotationDate"`
	ContractIdentification string `json:"contractIdentification"`
}

type AccountTransactions struct {
	Transactions TransactionList `json:"transactions"`
}

type TransactionList struct {
	Booked  []Transaction `json:"booked"`
	Pending []Transaction `json:"pending"`
}
