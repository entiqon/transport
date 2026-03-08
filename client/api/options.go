package api

import (
	"net/http"

	"github.com/entiqon/transport/auth"
)

// Option configures the API client.
type Option func(*api)

// WithHTTPClient sets a custom HTTP client.
func WithHTTPClient(c *http.Client) Option {
	return func(a *api) {
		a.http = c
	}
}

// WithAuth sets the authentication strategy.
func WithAuth(a auth.Auth) Option {
	return func(c *api) {
		c.auth = a
	}
}
