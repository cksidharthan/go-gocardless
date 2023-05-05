package consts

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	Summary    string `json:"summary"`
	Detail     string `json:"detail"`
	StatusCode int    `json:"status_code"`
	Reference  string `json:"reference"`
}

type Reference struct {
	Summary string `json:"summary"`
	Detail  string `json:"detail"`
}

func NewError(errResponse *http.Response) error {
	var newErr Error
	// print response body
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
