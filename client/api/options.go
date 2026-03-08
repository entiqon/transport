package api

import (
	"net/http"

	"github.com/entiqon/transport/auth"
)

// Option represents a configuration option for the API client.
type Option func(*api)

// WithHTTPClient configures the HTTP client used to perform requests.
func WithHTTPClient(client *http.Client) Option {
	return func(a *api) {
		a.http = client
	}
}

// WithAuth configures the authentication strategy applied to requests.
func WithAuth(strategy auth.Auth) Option {
	return func(a *api) {
		a.auth = strategy
	}
}
