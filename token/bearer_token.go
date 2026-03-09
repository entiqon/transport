package token

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/entiqon/transport/auth"
)

// BearerToken implements auth.Credential using the HTTP Authorization
// header with the Bearer authentication scheme.
//
// This strategy is commonly used with OAuth2 APIs and services
// that require a static or externally managed bearer token.
type bearerToken struct {
	token string
}

// NewBearerToken creates a new Bearer token authentication strategy.
func NewBearerToken(token string) auth.Credential {
	return &bearerToken{
		token: token,
	}
}

// Apply adds the Authorization header using the Bearer scheme.
func (b *bearerToken) Apply(_ context.Context, r *http.Request) error {
	if strings.TrimSpace(b.token) == "" {
		return fmt.Errorf("token is empty")
	}

	r.Header.Set("Authorization", "Bearer "+b.token)
	return nil
}
