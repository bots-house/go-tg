package tg

import "net/http"

// Client it's Telegram Bot API client.
type Client struct {
	token string
	doer  *http.Client
}

// ClientOption define type for custumize client.
type ClientOption func(*Client)

// WithDoer sets client http doer.
func WithDoer(doer *http.Client) ClientOption {
	return func(client *Client) {
		client.doer = doer
	}
}

// NewClient creates new Telegram Bot API client with provided token.
// Additional options can be specified using ClientOption.
func NewClient(token string, opts ...ClientOption) *Client {
	// create client with default values
	client := &Client{
		token: token,
		doer:  http.DefaultClient,
	}

	// customize client via options
	for _, opt := range opts {
		opt(client)
	}

	return client
}
