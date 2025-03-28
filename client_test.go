package gocardless_test

import (
	"testing"

	"github.com/cksidharthan/go-gocardless"
)

func getInvalidTestClient(t *testing.T) (*gocardless.Client, error) {
	t.Helper()

	return &gocardless.Client{
		HTTP:      gocardless.NewHTTPClient(gocardless.BaseURL, gocardless.APIVersion),
		SecretID:  "invalid",
		SecretKey: "invalid",
		Token: &gocardless.Token{
			Access:  "invalid",
			Refresh: "invalid",
		},
	}, nil
}
