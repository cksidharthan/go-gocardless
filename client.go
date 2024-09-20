package nordigen

// Client is the Nordigen client
type Client struct {
	HTTP      IHTTPClient
	SecretID  string
	SecretKey string
}

type Config struct {
	BaseURL    string
	APIVersion string
	SecretID   string `json:"secret_id"`
	SecretKey  string `json:"secret_key"`
	HTTP       *Client
}

// New creates a new Nordigen client
func New(config *Config) *Client {
	return &Client{
		HTTP:      NewHTTPClient(config.BaseURL, config.APIVersion),
		SecretID:  config.SecretID,
		SecretKey: config.SecretKey,
	}
}
