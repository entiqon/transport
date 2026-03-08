package auth

import (
	"context"
	"net/http"
)

// Auth defines a mechanism that can apply authentication
// to an outgoing HTTP request.
type Auth interface {
	Apply(ctx context.Context, req *http.Request) error
}
