package requisitions

import "time"

type ListRequestParams struct {
	Limit  int `url:"limit,omitempty" json:"limit,omitempty"`
	Offset int `url:"offset,omitempty" json:"offset,omitempty"`
}

type RequisitionRequestBody struct {
	Redirect          string `json:"redirect"`
	InstitutionID     string `json:"institution_id"`
	Agreement         string `json:"agreement"`
	Reference         string `json:"reference"`
	UserLanguage      string `json:"user_language"`
	SSN               string `json:"ssn"`
	AccountSelection  bool   `json:"account_selection"`
	RedirectImmediate bool   `json:"redirect_immediate"`
}

type Requisition struct {
	ID                string    `json:"id"`
	Created           time.Time `json:"created"`
	Redirect          string    `json:"redirect"`
	Status            string    `json:"status"`
	InstitutionID     string    `json:"institution_id"`
	Agreement         string    `json:"agreement"`
	Reference         string    `json:"reference"`
	Accounts          []string  `json:"accounts"`
	UserLanguage      string    `json:"user_language"`
	Link              string    `json:"link"`
	SSN               string    `json:"ssn"`
	AccountSelection  bool      `json:"account_selection"`
	RedirectImmediate bool      `json:"redirect_immediate"`
}

type Requisitions struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []Requisition `json:"results"`
}
