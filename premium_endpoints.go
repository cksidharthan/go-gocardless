package gocardless

import "context"

// ListPremiumTransactions retrieves transactions for an account by ID
func (c Client) ListPremiumTransactions(ctx context.Context, accountID string) (*PremiumTransactions, error) {
	var transactions PremiumTransactions
	endpointURL := AccountsTransactionPremiumPath + accountID + "/transactions"
	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(c.Token.Access), &transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}
