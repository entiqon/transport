package api

import (
	"net/http"

	"github.com/entiqon/transport/auth"
)

// Option configures the API client.
type Option func(*api)

// WithHTTPClient configures the HTTP client used to perform requests.
func WithHTTPClient(client *http.Client) Option {
	return func(a *api) {
		a.http = client
	}
}

// WithCredential configures the credential strategy applied to outgoing requests.
func WithCredential(credential auth.Credential) Option {
	return func(a *api) {
		a.credential = credential
	}
}
