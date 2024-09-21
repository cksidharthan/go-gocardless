package gocardless

import "time"

type PremiumTransactions struct {
	Transactions PremiumTransactionList `json:"transactions"`
}

type PremiumTransactionList struct {
	Booked  []PremiumTransaction `json:"booked"`
	Pending []PremiumTransaction `json:"pending"`
}

type PremiumTransaction struct {
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

	AdditionalInformation           string  `json:"additionalInformation"`
	AdditionalInformationStructured string  `json:"additionalInformationStructured"`
	BalanceAfterTransaction         Balance `json:"balanceAfterTransaction"`
	CheckID                         string  `json:"checkId"`
	CreditorID                      string  `json:"creditorId"`
	// CurrencyExchange                []string `json:"currencyExchange"`
	DebtorAgent                          string     `json:"debtorAgent"`
	EndToEndID                           string     `json:"endToEndId"`
	EntryReference                       string     `json:"entryReference"`
	MandateID                            string     `json:"mandateId"`
	MerchantCategoryCode                 string     `json:"merchantCategoryCode"`
	RemittanceInformationStructured      string     `json:"remittanceInformationStructured"`
	RemittanceInformationStructuredArray []string   `json:"remittanceInformationStructuredArray"`
	UltimateCollector                    string     `json:"ultimateCreditor"`
	UltimateDebtor                       string     `json:"ultimateDebtor"`
	Enrichment                           Enrichment `json:"enrichment"`
}

type Enrichment struct {
	DisplayName       string   `json:"displayName"`
	BranchDisplayName string   `json:"branchDisplayName"`
	Location          Location `json:"location"`
	URLs              URLs     `json:"urls"`
	TransactionType   string   `json:"transactionType"`
	PurposeCategory   []string `json:"purposeCategory"`
	PurposeCategoryID string   `json:"purposeCategoryID"`
}

type Location struct {
	Address    string  `json:"address"`
	City       string  `json:"city"`
	Region     string  `json:"region"`
	PostalCode string  `json:"postalCode"`
	Country    string  `json:"country"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lon"`
}

type URLs struct {
	Website string `json:"website"`
	Favicon string `json:"favicon"`
	Logo    string `json:"logo"`
}
