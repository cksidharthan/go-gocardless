package gocardless

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Client is the Nordigen client
type Client struct {
	mu        sync.Mutex
	Logger    *log.Logger
	HTTP      IHTTPClient
	SecretID  string
	SecretKey string
	Token     *Token
}

type Config struct {
	Logger       *log.Logger
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

	if config.Logger != nil {
		client.Logger = config.Logger
	} else {
		client.Logger = slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelInfo)
	}

	token, err := client.NewToken(context.Background())
	if err != nil {
		client.Logger.Printf("failed to get token: %v\n", err)
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	client.Token = token

	if config.TokenRefresh {
		go refreshGocardlessToken(client)
	}

	return client, nil
}

// refreshGocardlessToken refreshes the access token and the refresh token
func refreshGocardlessToken(client *Client) {
	client.Logger.Println("refreshGocardlessToken started")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGTERM, syscall.SIGINT)

	for {
		select {
		case <-sigterm:
			fmt.Println("Received termination signal, exiting refreshGocardlessToken")

			return
		case <-time.After(1 * time.Minute):
			if time.Unix(int64(client.Token.AccessExpires), 0).Before(time.Now().Add(1 * time.Minute)) {
				newToken, err := client.RefreshToken(context.Background(), client.Token.Refresh)
				if err != nil {
					client.Logger.Printf("failed to refresh access token: %v\n", err)
					continue
				}

				client.mu.Lock()
				client.Token.Access = newToken.Access
				client.Token.AccessExpires = newToken.AccessExpires
				client.mu.Unlock()
			}

			if time.Unix(int64(client.Token.RefreshExpires), 0).Before(time.Now().Add(1 * time.Minute)) {
				newToken, err := client.NewToken(context.Background())
				if err != nil {
					client.Logger.Printf("failed to create new token: %v\n", err)
					continue
				}

				client.mu.Lock()
				client.Token = newToken
				client.mu.Unlock()
			}
		}
	}
}
