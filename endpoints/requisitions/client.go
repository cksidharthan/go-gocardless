package requisitions

import "github.com/weportfolio/go-nordigen/http"

type Client struct {
	HTTP *http.Client
}

func New(secretID, secretKey string) *Client {
	return &Client{
		HTTP: http.New(secretID, secretKey),
	}
}
