package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/weportfolio/go-nordigen/consts"
	"net/http"
)

const (
	NordigenBaseURL = "https://ob.nordigen.com/api"
	APIVersion      = "v2"
)

//go:generate mockgen -source=http-client.go -destination=mocks/mock_client.go -package=mocks -build_flags=-mod=mod
type IClient interface {
	Get(ctx context.Context, path string, headers map[string]string, response interface{}) error
	Post(ctx context.Context, path string, headers map[string]string, body interface{}, response interface{}) error
	Put(ctx context.Context, path string, headers map[string]string, body interface{}, response interface{}) error
	Delete(ctx context.Context, path string, headers map[string]string, response interface{}) error
}

type Client struct {
	BaseURL      string
	APIVersion   string
	APISecretID  string
	APISecretKey string
}

func New(secretID, secretKey string) *Client {
	return &Client{
		BaseURL:      NordigenBaseURL,
		APIVersion:   APIVersion,
		APISecretID:  secretID,
		APISecretKey: secretKey,
	}
}

// Get is a wrapper around request that performs a GET request
func (c *Client) Get(ctx context.Context, path string, headers map[string]string, response interface{}) error {
	return c.request(ctx, http.MethodGet, path, headers, nil, response)
}

// Post is a wrapper around request that performs a POST request
func (c *Client) Post(ctx context.Context, path string, headers map[string]string, body interface{}, response interface{}) error {
	return c.request(ctx, http.MethodPost, path, headers, body, response)
}

// Put is a wrapper around request that performs a PUT request
func (c *Client) Put(ctx context.Context, path string, headers map[string]string, body interface{}, response interface{}) error {
	return c.request(ctx, http.MethodPut, path, headers, body, response)
}

// Delete is a wrapper around request that performs a DELETE request
func (c *Client) Delete(ctx context.Context, path string, headers map[string]string, response interface{}) error {
	return c.request(ctx, http.MethodDelete, path, headers, nil, response)
}

// request is a wrapper around http.Client.Do that performs a request and decodes the response into the response interface
func (c *Client) request(ctx context.Context, method, path string, headers map[string]string, body interface{}, response interface{}) error {
	var bytesBody []byte
	var err error
	if body != nil {
		bytesBody, err = convertToBytes(body)
		if err != nil {
			return fmt.Errorf("failed to convert body to bytes: %w", err)
		}
	}

	req, err := c.newRequest(ctx, method, path, headers, bytesBody)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	if resp.StatusCode > 300 {
		return consts.NewError(resp)
	}

	if response != nil {
		return decodeResponse(resp, response)
	}

	return nil
}

// newRequest creates a new http.Request with the given method, path, headers and body
func (c *Client) newRequest(ctx context.Context, method, path string, headers map[string]string, body []byte) (*http.Request, error) {
	url := c.BaseURL + "/" + c.APIVersion + path

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
	} else {
		req, err = http.NewRequestWithContext(ctx, method, url, http.NoBody)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	return req, nil
}

func decodeResponse(resp *http.Response, response interface{}) error {
	err := json.NewDecoder(resp.Body).Decode(response)
	if err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

func convertToBytes(requestBody interface{}) (body []byte, err error) {
	body, err = json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request body: %w", err)
	}

	return body, nil
}
