package accounts

import (
	"encoding/json"
	"fmt"
	"io"
)

// Get retrieves an account by ID
func (a Client) Get(accountID string) (account *Account, err error) {
	path := fmt.Sprintf("/accounts/%s", accountID)
	response, err := a.HTTP.Get(path, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to perform request: %w", err)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Decode response into account
	err = json.Unmarshal(body, &account)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return account, nil
}
