package gocardless_test

import (
	"github.com/cksidharthan/go-gocardless"
	"log"
	"os"
	"testing"
)

var (
	testClient    *gocardless.Client
	testAccountID string
)

func TestMain(m *testing.M) {
	var err error

	gocardlessSecretID := os.Getenv("GOCARDLESS_SECRET_ID")
	if gocardlessSecretID == "" {
		log.Fatal("GOCARDLESS_SECRET is not set")
	}

	gocardlessSecretKey := os.Getenv("GOCARDLESS_SECRET_KEY")
	if gocardlessSecretKey == "" {
		log.Fatal("GOCARDLESS_SECRET_KEY is not set")
	}

	testClient, err = gocardless.New(&gocardless.Config{
		BaseURL:    gocardless.BaseURL,
		APIVersion: gocardless.APIVersion,
		SecretID:   gocardlessSecretID,
		SecretKey:  gocardlessSecretKey,
	})
	if err != nil {
		log.Fatalf("failed to create client: %v\n", err)
	}

	// Run tests
	code := m.Run()

	// Exit
	os.Exit(code)
}
