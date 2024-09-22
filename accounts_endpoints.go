package gocardless

import (
	"context"
	"fmt"
	"net/url"
	"time"
)

// GetAccount retrieves an account by ID
func (c Client) GetAccount(ctx context.Context, accountID string) (*Account, error) {
	var account Account
	endpointURL := AccountsPath + accountID
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(c.Token.Access), &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetAccountBalances retrieves balances for an account by ID
func (c Client) GetAccountBalances(ctx context.Context, accountID string) (*Balances, error) {
	var balances Balances
	endpointURL := AccountsPath + accountID + "/balances"
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(c.Token.Access), &balances)
	if err != nil {
		return nil, err
	}

	return &balances, nil
}

// GetAccountDetails retrieves details for an account by ID
func (c Client) GetAccountDetails(ctx context.Context, accountID string) (*Details, error) {
	var details Details
	endpointURL := AccountsPath + accountID + "/details"
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(c.Token.Access), &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// GetAccountTransactions retrieves transactions for an account by ID
// dateFrom and dateTo are optional parameters.
func (c Client) GetAccountTransactions(ctx context.Context, accountID string, dateFrom, dateTo *time.Time) (*Transactions, error) {
	var transactions Transactions
	endpointURL := AccountsPath + accountID + "/transactions"

	// Build query parameters
	u, err := url.Parse(endpointURL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse URL: %w", err)
	}

	q := u.Query()
	if dateFrom != nil {
		q.Set("date_from", dateFrom.Format(time.RFC3339))
	}
	if dateTo != nil {
		q.Set("date_to", dateTo.Format(time.RFC3339))
	}
	u.RawQuery = q.Encode()

	err = c.HTTP.Get(ctx, u.String(), RequestHeadersWithAuth(c.Token.Access), &transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}
