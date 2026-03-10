package auth

import "context"

// Refreshable indicates the provider can force a credential refresh.
type Refreshable interface {
	// Refresh forces the provider to obtain a new credential.
	Refresh(ctx context.Context) error
}
