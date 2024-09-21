package gocardless

import (
	"context"
)

// NewToken creates a new token
func (c Client) NewToken(ctx context.Context) (*Token, error) {
	var token Token
	accessCreds := map[string]string{
		"secret_id":  c.SecretID,
		"secret_key": c.SecretKey,
	}

	err := c.HTTP.Post(ctx, TokenNewPath, RequestHeaders(), accessCreds, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// RefreshToken refreshes a token
func (c Client) RefreshToken(ctx context.Context, refreshToken string) (*Token, error) {
	var token Token
	refreshCreds := map[string]string{
		"refresh": refreshToken,
	}

	err := c.HTTP.Post(ctx, TokenRefreshPath, RequestHeaders(), refreshCreds, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
