package credential

import (
	"context"
	"errors"
	"net/http"

	"github.com/entiqon/transport/auth"
)

// AccessToken represents a header-based token authentication strategy.
//
// It injects a token into the HTTP request using a configurable header.
// This pattern is commonly used by APIs such as Shopify and other
// header-token based authentication systems.
type accessToken struct {
	key   string
	value string
}

// AccessToken creates a new AccessToken authentication strategy.
func AccessToken(header, token string) auth.Credential {
	return &accessToken{
		key:   header,
		value: token,
	}
}

// Apply injects the access token into the request header.
func (a *accessToken) Apply(_ context.Context, r *http.Request) error {
	if a.key == "" {
		return errors.New("api key missing")
	}

	r.Header.Set(a.key, a.value)
	return nil
}
