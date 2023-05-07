package nordigen

import "time"

type AgreementRequestBody struct {
	InstitutionID      string   `json:"institution_id"`
	MaxHistoricalDays  string   `json:"max_historical_days"`
	AccessValidForDays string   `json:"access_valid_for_days"`
	AccessScope        []string `json:"access_scope"`
}

type Agreement struct {
	ID                 string    `json:"id"`
	Created            time.Time `json:"created"`
	InstitutionID      string    `json:"institution_id"`
	MaxHistoricalDays  int       `json:"max_historical_days"`
	AccessValidForDays int       `json:"access_valid_for_days"`
	AccessScope        []string  `json:"access_scope"`
	Accepted           time.Time `json:"accepted"`
}

type ListAgreementsParams struct {
	Limit  int `url:"limit,omitempty" json:"limit,omitempty"`
	Offset int `url:"offset,omitempty" json:"offset,omitempty"`
}

type Agreements struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Agreement `json:"results"`
}

type UpdateRequestBody struct {
	UserAgent string `json:"user_agent"`
	IPAddress string `json:"ip_address"`
}
