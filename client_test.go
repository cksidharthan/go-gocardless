package gocardless_test

import (
	"fmt"
	"os"
	"testing"

	gocardless "github.com/cksidharthan/go-gocardless"
)

func getTestClient(t *testing.T) (*gocardless.Client, error) {
	t.Helper()

	secretID := os.Getenv("GOCARDLESS_SECRET_ID")
	secretKey := os.Getenv("GOCARDLESS_SECRET_KEY")
	if secretID == "" || secretKey == "" {
		fmt.Println("GOCARDLESS_SECRET_ID or GOCARDLESS_SECRET_KEY is not set")
	}

	return gocardless.New(
		&gocardless.Config{
			BaseURL:    gocardless.NordigenBaseURL,
			APIVersion: gocardless.APIVersion,
			SecretID:   secretID,
			SecretKey:  secretKey,
		},
	)
}

func getInvalidTestClient(t *testing.T) (*gocardless.Client, error) {
	t.Helper()

	return gocardless.New(
		&gocardless.Config{
			BaseURL:    gocardless.NordigenBaseURL,
			APIVersion: gocardless.APIVersion,
			SecretID:   "invalid",
			SecretKey:  "invalid",
		},
	)
}
