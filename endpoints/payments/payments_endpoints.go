package payments

import (
	"context"
	"strconv"

	"github.com/weportfolio/go-nordigen/consts"
)

// List returns a list of payments
func (c Client) List(ctx context.Context, token string, limit, offset int) (*Payments, error) {
	var response Payments

	params := map[string]string{
		"limit":  strconv.Itoa(limit),
		"offset": strconv.Itoa(offset),
	}

	err := c.HTTP.Get(ctx, consts.BuildQueryURL(consts.PaymentsPath, params), consts.RequestHeadersWithAuth(token), &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
