package nordigen_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/weportfolio/go-nordigen"
)

func getTestClient(t *testing.T) *nordigen.Client {
	t.Helper()

	secretID := os.Getenv("NORDIGEN_SECRET_ID")
	secretKey := os.Getenv("NORDIGEN_SECRET_KEY")
	if secretID == "" || secretKey == "" {
		fmt.Println("NORDIGEN_SECRET_ID or NORDIGEN_SECRET_KEY is not set")
	}

	return nordigen.New(
		&nordigen.Config{
			BaseURL:    nordigen.NordigenBaseURL,
			APIVersion: nordigen.APIVersion,
			SecretID:   secretID,
			SecretKey:  secretKey,
		},
	)
}

func getInvalidTestClient(t *testing.T) *nordigen.Client {
	t.Helper()

	return nordigen.New(
		&nordigen.Config{
			BaseURL:    nordigen.NordigenBaseURL,
			APIVersion: nordigen.APIVersion,
			SecretID:   "invalid",
			SecretKey:  "invalid",
		},
	)
}
