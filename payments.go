package nordigen

type Payment struct {
	PaymentID        string         `json:"payment_id"`
	PaymentStatus    string         `json:"payment_status"`
	PaymentProduct   string         `json:"payment_product"`
	PaymentType      string         `json:"payment_type"`
	Redirect         string         `json:"redirect"`
	Description      string         `json:"description"`
	CustomPaymentID  string         `json:"custom_payment_id"`
	CreditorAccount  string         `json:"creditor_account"`
	CreditorObject   PaymentAccount `json:"creditor_object"`
	DebtorObject     PaymentAccount `json:"debtor_object"`
	InstructedAmount Amount         `json:"instructed_amount"`
}

type Payments struct {
	Code     int       `json:"code"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Payment `json:"results"`
}

type PaymentAccount struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	Type           string `json:"type"`
	Account        string `json:"account"`
	Currency       string `json:"currency"`
	AddressCountry string `json:"address_country"`
	InstitutionID  string `json:"institution_id"`
	Agent          string `json:"agent"`
	AgentName      string `json:"agent_name"`
	AddressStreet  string `json:"address_street"`
	PostCode       string `json:"post_code"`
	TypeNumber     string `json:"type_number"`
}
