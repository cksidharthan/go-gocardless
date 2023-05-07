package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/weportfolio/go-nordigen"
	"github.com/weportfolio/go-nordigen/consts"
)

func GetTestClient(t *testing.T) *nordigen.Client {
	t.Helper()

	secretID := os.Getenv("NORDIGEN_SECRET_ID")
	secretKey := os.Getenv("NORDIGEN_SECRET_KEY")
	if secretID == "" || secretKey == "" {
		fmt.Println("NORDIGEN_SECRET_ID or NORDIGEN_SECRET_KEY is not set")
	}

	return nordigen.New(
		&nordigen.Config{
			BaseURL:    consts.NordigenBaseURL,
			APIVersion: consts.APIVersion,
			SecretID:   secretID,
			SecretKey:  secretKey,
		},
	)
}

func GetInvalidTestClient(t *testing.T) *nordigen.Client {
	t.Helper()

	return nordigen.New(
		&nordigen.Config{
			BaseURL:    consts.NordigenBaseURL,
			APIVersion: consts.APIVersion,
			SecretID:   "invalid",
			SecretKey:  "invalid",
		},
	)
}
