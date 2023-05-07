package accounts

import (
	"context"
	"fmt"

	"github.com/weportfolio/go-nordigen/consts"
)

// GetAccount retrieves an account by ID
func (c Client) GetAccount(ctx context.Context, token string, accountID string) (*Account, error) {
	var account Account
	endpointURL := fmt.Sprintf("/accounts/%s", accountID)
	err := c.HTTP.Get(ctx, endpointURL, consts.RequestHeadersWithAuth(token), &account)
	if err != nil {
		return nil, err
	}
	return &account, nil
}

// GetBalances retrieves balances for an account by ID
func (c Client) GetBalances(ctx context.Context, token string, accountID string) (*Balances, error) {
	var balances Balances
	endpointURL := fmt.Sprintf("/accounts/%s/balances", accountID)
	err := c.HTTP.Get(ctx, endpointURL, consts.RequestHeadersWithAuth(token), &balances)
	if err != nil {
		return nil, err
	}

	return &balances, nil
}

func (c Client) GetDetails(ctx context.Context, token string, accountID string) (*Details, error) {
	var details Details
	endpointURL := fmt.Sprintf("/accounts/%s/details", accountID)
	err := c.HTTP.Get(ctx, endpointURL, consts.RequestHeadersWithAuth(token), &details)
	if err != nil {
		return nil, err
	}

	return &details, nil
}

func (c Client) GetTransactions(ctx context.Context, token string, accountID string) (*Transactions, error) {
	var transactions Transactions
	endpointURL := fmt.Sprintf("/accounts/%s/transactions", accountID)
	err := c.HTTP.Get(ctx, endpointURL, consts.RequestHeadersWithAuth(token), &transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}
