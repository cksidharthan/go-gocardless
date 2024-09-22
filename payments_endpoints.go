package gocardless

import (
	"context"
	"strconv"
)

// ListPayments returns a list of payments
func (c Client) ListPayments(ctx context.Context, limit, offset int) (*Payments, error) {
	var response Payments

	params := map[string]string{
		"limit":  strconv.Itoa(limit),
		"offset": strconv.Itoa(offset),
	}

	err := c.HTTP.Get(ctx, BuildQueryURL(PaymentsPath, params), RequestHeadersWithAuth(c.Token.Access), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
