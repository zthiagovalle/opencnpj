package opencnpj

import (
	"net/http"
	"time"
)

// Option is a function that configures a Client.
type Option func(*Client)

// WithTimeout sets the timeout for requests.
func WithTimeout(d time.Duration) Option {
	return func(c *Client) {
		c.httpClient.Timeout = d
	}
}

// WithBaseURL sets the base URL for the API.
func WithBaseURL(url string) Option {
	return func(c *Client) {
		c.baseURL = url
	}
}

// WithHTTPClient sets the HTTP client for the client.
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
