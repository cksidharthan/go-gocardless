package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	NordigenBaseURL = "https://ob.nordigen.com/api"
	APIVersion      = "v2"
)

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

func (c *Client) Get(path string, params map[string]string, response interface{}) error {
	return c.request(http.MethodGet, path, params, nil, response)
}

func (c *Client) Post(path string, params map[string]string, body []byte, response interface{}) error {
	return c.request(http.MethodPost, path, params, body, response)
}

func (c *Client) Put(path string, params map[string]string, body []byte, response interface{}) error {
	return c.request(http.MethodPut, path, params, body, response)
}

func (c *Client) Delete(path string, params map[string]string, response interface{}) error {
	return c.request(http.MethodDelete, path, params, nil, response)
}

func (c *Client) request(method, path string, params map[string]string, body []byte, response interface{}) error {
	req, err := c.newRequest(method, path, params, body)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("failed to close response body: %s\n", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed with status code %d", resp.StatusCode)
	}

	if response != nil {
		err = json.NewDecoder(resp.Body).Decode(response)
		if err != nil {
			return fmt.Errorf("failed to decode response: %w", err)
		}
	}

	return nil
}

func (c *Client) newRequest(method, path string, params map[string]string, body []byte) (*http.Request, error) {
	url := c.BaseURL + "/" + c.APIVersion + path

	var req *http.Request
	var err error

	if body != nil {
		req, err = http.NewRequest(method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
	} else {
		req, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to create request: %w", err)
		}
	}

	req.SetBasicAuth(c.APISecretID, c.APISecretKey)

	q := req.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}
