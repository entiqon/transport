package auth

import (
	"context"

	"github.com/entiqon/transport/config"
)

// Provider resolves authentication configuration into a Credential.
//
// A Provider interprets authentication configuration and returns a
// Credential that can authenticate outgoing HTTP requests.
//
// Implementations may perform operations such as token retrieval,
// token refresh, or credential construction before returning the
// resulting Credential.
//
// Providers do not mutate HTTP requests directly. Request mutation
// is the responsibility of the returned Credential.
type Provider interface {
	// Resolve returns a Credential based on the given authentication
	// configuration.
	Resolve(ctx context.Context, cfg config.Auth) (Credential, error)
}
