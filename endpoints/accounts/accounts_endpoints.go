package accounts

import (
	"context"
	"fmt"
)

// Get retrieves an account by ID
func (a Client) Get(ctx context.Context, accountID string) (account *Account, err error) {
	path := fmt.Sprintf("/accounts/%s", accountID)
	err = a.HTTP.Get(ctx, path, nil, account)
	return account, nil
}
