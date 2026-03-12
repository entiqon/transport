package api

import (
	"net/http"
	"strings"

	"github.com/entiqon/transport/auth"
	"github.com/entiqon/transport/config"
)

// Option configures the API client.
type Option func(*client)

// WithHTTPClient configures the HTTP client used to perform requests.
func WithHTTPClient(http *http.Client) Option {
	return func(a *client) {
		a.http = http
	}
}

// WithCredential configures the credential strategy applied to outgoing requests.
func WithCredential(credential auth.Credential) Option {
	return func(a *client) {
		a.credential = credential
	}
}

// WithAuthProvider configures the authentication provider and its
// configuration used to resolve credentials for outgoing requests.
func WithAuthProvider(
	provider auth.Provider,
	cfg config.Auth,
) Option {
	return func(a *client) {
		a.provider = provider
		a.config = cfg
	}
}

// WithBasePath configures a base path prefix applied to all requests
// executed by the client.
//
// The base path represents an API namespace such as "api" or "v1".
// Leading and trailing slashes are automatically removed.
//
// For example:
//
//	WithBasePath("api")
//
// combined with a request path "users" results in:
//
//	/api/users
func WithBasePath(basePath string) Option {
	return func(c *client) {
		c.basePath = strings.Trim(basePath, "/")
	}
}

// WithVersion sets the API version segment inserted between
// the base URL and request path when building requests.
func WithVersion(version string) Option {
	return func(c *client) {
		c.version = strings.Trim(version, "/")
	}
}
