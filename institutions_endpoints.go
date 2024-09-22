package gocardless

import (
	"context"
)

// ListInstitutions returns a list of institutions
func (c Client) ListInstitutions(ctx context.Context, country string, paymentsEnabled bool) ([]Institution, error) {
	endpointURL := InstitutionsPath + "?country=" + country + "&payments_enabled=" + boolToString(paymentsEnabled)
	var institutions []Institution

	err := c.HTTP.Get(ctx, endpointURL, RequestHeadersWithAuth(c.Token.Access), &institutions)
	if err != nil {
		return nil, err
	}

	return institutions, nil
}

// FetchInstitution returns an institution
func (c Client) FetchInstitution(ctx context.Context, id string) (*Institution, error) {
	var institution Institution
	err := c.HTTP.Get(ctx, InstitutionsPath+id, RequestHeadersWithAuth(c.Token.Access), &institution)
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
