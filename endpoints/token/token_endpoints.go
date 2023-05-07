package token

import (
	"context"

	"github.com/weportfolio/go-nordigen/consts"
)

// New creates a new token
func (c Client) New(ctx context.Context) (*Token, error) {
	var token Token
	accessCreds := map[string]string{
		"secret_id":  c.SecretID,
		"secret_key": c.SecretKey,
	}

	err := c.HTTP.Post(ctx, consts.TokenNewPath, consts.RequestHeaders(), accessCreds, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// Refresh refreshes a token
func (c Client) Refresh(ctx context.Context, refreshToken string) (*Token, error) {
	var token Token
	refreshCreds := map[string]string{
		"refresh": refreshToken,
	}

	err := c.HTTP.Post(ctx, consts.TokenRefreshPath, consts.RequestHeaders(), refreshCreds, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}
