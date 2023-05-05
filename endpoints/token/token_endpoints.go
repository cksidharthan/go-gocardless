package token

import (
	"encoding/json"
	"fmt"
	"io"
)

func (c Client) New() (*Token, error) {
	var token Token
	accessCreds := map[string]string{
		"secret_id":  c.HTTP.APISecretID,
		"secret_key": c.HTTP.APISecretKey,
	}

	body, err := convertToBytes(accessCreds)
	if err != nil {
		return nil, fmt.Errorf("failed to convert accessCreds to bytes: %w", err)
	}

	response, err := c.HTTP.Post("/token/new/", nil, body)
	if err != nil {
		return nil, err
	}

	body, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	defer response.Body.Close()

	// Decode response into token
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response body: %w", err)
	}

	return &token, nil
}

func convertToBytes(params map[string]string) (body []byte, err error) {
	body, err = json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal params: %w", err)
	}

	return body, nil
}
