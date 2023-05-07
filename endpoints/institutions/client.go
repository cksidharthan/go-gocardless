package institutions

import "github.com/weportfolio/go-nordigen/http"

type Client struct {
	HTTP *http.Client
}

func New(baseURL, apiVersion string) *Client {
	return &Client{
		HTTP: http.New(baseURL, apiVersion),
	}
}
