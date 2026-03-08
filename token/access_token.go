package token

import (
	"context"
	"net/http"

	"github.com/entiqon/transport/auth"
)

// AccessToken represents a header-based token authentication strategy.
//
// It injects a token into the HTTP request using a configurable header.
// This pattern is commonly used by APIs such as Shopify and other
// header-token based authentication systems.
type accessToken struct {
	Header string
	Token  string
}

// NewAccessToken creates a new AccessToken authentication strategy.
func NewAccessToken(header, token string) auth.Auth {
	return &accessToken{
		Header: header,
		Token:  token,
	}
}

// Apply injects the access token into the request header.
func (a *accessToken) Apply(_ context.Context, r *http.Request) error {
	r.Header.Set(a.Header, a.Token)
	return nil
}
