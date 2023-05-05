package nordigen

import (
	"github.com/weportfolio/go-nordigen/endpoints/accounts"
	"github.com/weportfolio/go-nordigen/endpoints/token"
)

// Client is the Nordigen client
type Client struct {
	accounts *accounts.Client
	token    *token.Client
}

// Accounts returns the accounts client
func (c *Client) Accounts() *accounts.Client {
	return c.accounts
}

// Token returns the token client
func (c *Client) Token() *token.Client {
	return c.token
}

// New creates a new Nordigen client
func New(secretID, secretKey string) *Client {
	return &Client{
		accounts: accounts.New(secretID, secretKey),
		token:    token.New(secretID, secretKey),
	}
}
