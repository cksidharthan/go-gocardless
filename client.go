package nordigen

import (
	"github.com/weportfolio/go-nordigen/endpoints/accounts"
	"github.com/weportfolio/go-nordigen/endpoints/agreements"
	"github.com/weportfolio/go-nordigen/endpoints/institutions"
	"github.com/weportfolio/go-nordigen/endpoints/payments"
	"github.com/weportfolio/go-nordigen/endpoints/premium"
	"github.com/weportfolio/go-nordigen/endpoints/requisitions"
	"github.com/weportfolio/go-nordigen/endpoints/token"
)

// Client is the Nordigen client
type Client struct {
	accounts     *accounts.Client
	token        *token.Client
	institutions *institutions.Client
	agreements   *agreements.Client
	requisitions *requisitions.Client
	payments     *payments.Client
	premium      *premium.Client
}

// Accounts returns the accounts client
func (c *Client) Accounts() *accounts.Client {
	return c.accounts
}

// Token returns the token client
func (c *Client) Token() *token.Client {
	return c.token
}

// Institutions returns the institutions client
func (c *Client) Institutions() *institutions.Client {
	return c.institutions
}

// Agreements returns the agreements client
func (c *Client) Agreements() *agreements.Client {
	return c.agreements
}

// Requisitions returns the requisitions client
func (c *Client) Requisitions() *requisitions.Client {
	return c.requisitions
}

// Payments returns the payments client
func (c *Client) Payments() *payments.Client {
	return c.payments
}

// Premium returns the premium client
func (c *Client) Premium() *premium.Client {
	return c.premium
}

type Config struct {
	BaseURL    string
	APIVersion string
	SecretID   string
	SecretKey  string
}

// New creates a new Nordigen client
func New(config *Config) *Client {
	return &Client{
		accounts:     accounts.New(config.BaseURL, config.APIVersion),
		token:        token.New(config.BaseURL, config.APIVersion, config.SecretID, config.SecretKey),
		institutions: institutions.New(config.BaseURL, config.APIVersion),
		agreements:   agreements.New(config.BaseURL, config.APIVersion),
		requisitions: requisitions.New(config.BaseURL, config.APIVersion),
		payments:     payments.New(config.BaseURL, config.APIVersion),
		premium:      premium.New(config.BaseURL, config.APIVersion),
	}
}
