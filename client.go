package nordigen

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Client is the Nordigen client
type Client struct {
	HTTP      IHTTPClient
	SecretID  string
	SecretKey string
	Token     *Token
}

type Config struct {
	BaseURL      string
	APIVersion   string
	SecretID     string `json:"secret_id"`
	SecretKey    string `json:"secret_key"`
	TokenRefresh bool
	HTTP         *Client
}

// New creates a new Gocardless client
func New(config *Config) (*Client, error) {
	client := &Client{
		HTTP:      NewHTTPClient(config.BaseURL, config.APIVersion),
		SecretID:  config.SecretID,
		SecretKey: config.SecretKey,
	}

	token, err := client.NewToken(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	client.Token = token

	if config.TokenRefresh {
		go refreshGocardlessToken(client)
	}

	return client, nil
}

// refreshGocardlessToken refreshes the access token and the refresh token
func refreshGocardlessToken(client *Client) error {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-sigterm:
			fmt.Println("Received termination signal, exiting refreshGocardlessToken")
			return nil
		case <-time.After(1 * time.Minute):
			if time.Unix(int64(client.Token.AccessExpires), 0).Before(time.Now().Add(1 * time.Minute)) {
				newToken, err := client.RefreshToken(context.Background(), client.Token.Refresh)
				if err != nil {
					fmt.Printf("failed to refresh access token: %v\n", err)
					continue
				}

				client.Token.Access = newToken.Access
				client.Token.AccessExpires = newToken.AccessExpires
			}

			if time.Unix(int64(client.Token.RefreshExpires), 0).Before(time.Now().Add(1 * time.Minute)) {
				newToken, err := client.NewToken(context.Background())
				if err != nil {
					fmt.Printf("failed to create new token: %v\n", err)
					continue
				}

				client.Token = newToken
			}
		}
	}
}
