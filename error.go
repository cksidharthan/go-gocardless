package gocardless

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Summary    string `json:"summary"`
	Detail     string `json:"detail"`
	StatusCode int    `json:"status_code"`
	Type       string `json:"type"`
}

func NewError(errResponse *http.Response) error {
	var newErr Error
	err := json.NewDecoder(errResponse.Body).Decode(&newErr)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	if newErr.Summary == "" {
		newErr.Summary = errResponse.Status
	}

	if newErr.Detail == "" {
		newErr.Detail = errResponse.Status
	}

	if newErr.StatusCode == 0 {
		newErr.StatusCode = errResponse.StatusCode
	}

	return &newErr
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d - %s: %s", e.StatusCode, e.Summary, e.Detail)
}

func ExtractError(err error) *Error {
	return err.(*Error)
}
