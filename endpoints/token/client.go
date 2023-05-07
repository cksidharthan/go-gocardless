package token

import "github.com/weportfolio/go-nordigen/http"

type Client struct {
	HTTP      *http.Client
	SecretID  string
	SecretKey string
}

func New(baseURL, apiVersion, secretID, secretKey string) *Client {
	return &Client{
		HTTP:      http.New(baseURL, apiVersion),
		SecretID:  secretID,
		SecretKey: secretKey,
	}
}
