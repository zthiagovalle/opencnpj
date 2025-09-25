package opencnpj

import (
	"net/http"
	"time"
)

// Client is a client for the OpenCNPJ API.
type Client struct {
	baseURL    string
	httpClient *http.Client
}

// NewClient creates a new OpenCNPJ client.
func NewClient(opts ...Option) *Client {
	client := &Client{
		baseURL: apiURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, opt := range opts {
		opt(client)
	}

	return client
}
