package gocardless

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Summary    string    `json:"summary"`
	Detail     string    `json:"detail"`
	StatusCode int       `json:"status_code"`
	Reference  Reference `json:"reference"`
	Type       string    `json:"type"`
}

type Reference struct {
	Summary string `json:"summary"`
	Detail  string `json:"detail"`
}

func NewError(errResponse *http.Response) error {
	var newErr Error
	err := json.NewDecoder(errResponse.Body).Decode(&newErr)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return &newErr
}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Summary, e.Detail)
}

func ExtractError(err error) *Error {
	return err.(*Error)
}
