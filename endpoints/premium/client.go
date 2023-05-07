package premium

import "github.com/weportfolio/go-nordigen/http"

type Client struct {
	HTTP http.IClient
}

func New(baseURL, apiVersion string) *Client {
	return &Client{
		HTTP: http.New(baseURL, apiVersion),
	}
}
