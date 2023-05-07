package nordigen

import (
	"context"
)

// GetAccount retrieves an account by ID
func (c Client) GetAccount(ctx context.Context, token string, accountID string) (*Account, error) {
	var account Account
	endpointURL := AccountsPath + accountID
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(token), &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetAccountBalances retrieves balances for an account by ID
func (c Client) GetAccountBalances(ctx context.Context, token string, accountID string) (*Balances, error) {
	var balances Balances
	endpointURL := AccountsPath + accountID + "/balances"
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(token), &balances)
	if err != nil {
		return nil, err
	}

	return &balances, nil
}

// GetAccountDetails retrieves details for an account by ID
func (c Client) GetAccountDetails(ctx context.Context, token string, accountID string) (*Details, error) {
	var details Details
	endpointURL := AccountsPath + accountID + "/details"
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(token), &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

// GetAccountTransactions retrieves transactions for an account by ID
func (c Client) GetAccountTransactions(ctx context.Context, token string, accountID string) (*Transactions, error) {
	var transactions Transactions
	endpointURL := AccountsPath + accountID + "/transactions"
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(token), &transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}
