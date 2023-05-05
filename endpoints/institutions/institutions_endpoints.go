package institutions

import (
	"context"
	"github.com/weportfolio/go-nordigen/consts"
)

// List returns a list of institutions
func (c *Client) List(ctx context.Context, token string, country string, paymentsEnabled bool) ([]Institution, error) {
	endpointURL := consts.InstitutionsPath + "?country=" + country + "&payments_enabled=" + boolToString(paymentsEnabled)
	var institutions []Institution

	err := c.HTTP.Get(ctx, endpointURL, consts.RequestHeadersWithAuth(token), &institutions)
	if err != nil {
		return nil, err
	}

	return institutions, nil
}

// Get returns an institution
func (c *Client) Get(ctx context.Context, token, id string) (*Institution, error) {
	var institution Institution
	err := c.HTTP.Get(ctx, consts.InstitutionsPath+"/"+id, consts.RequestHeadersWithAuth(token), &institution)
	if err != nil {
		return nil, err
	}

	return &institution, nil
}

// boolToString converts a bool to a string
func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
