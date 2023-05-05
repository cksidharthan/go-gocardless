package accounts

type Account struct {
	ID            string `json:"id"`
	Created       string `json:"created"`
	LastAccessed  string `json:"last_accessed"`
	IBAN          string `json:"iban"`
	InstitutionID string `json:"institution_id"`
	Status        string `json:"status"`
	OwnerName     string `json:"owner_name"`
}

type Error struct {
	Summary    string `json:"summary"`
	Detail     string `json:"detail"`
	StatusCode int    `json:"status_code"`
}
