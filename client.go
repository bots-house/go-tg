package tg

import "net/http"


// Doer define interface for mock client calls.
type Doer interface {
	// Do HTTP request.
	Do(r *http.Request) (*http.Response, error)
}

// ClientOption define type for customize client.
type ClientOption func(*Client)

// WithDoer sets client http doer.
func WithDoer(doer Doer) ClientOption {
	return func(client *Client) {
		client.doer = doer
	}
}

// Client it's Telegram Bot API client.
type Client struct {
	token string
	doer  Doer
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
