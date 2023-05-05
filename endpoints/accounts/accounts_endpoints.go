package accounts

import (
	"fmt"
)

// Get retrieves an account by ID
func (a Client) Get(accountID string) (account *Account, err error) {
	path := fmt.Sprintf("/accounts/%s", accountID)
	err = a.HTTP.Get(path, nil, account)
	return account, nil
}
