package auth

import (
	"context"
	"net/http"
)

// Auth represents a strategy capable of applying authentication
// information to an outgoing HTTP request.
//
// Implementations typically modify request headers (for example,
// Authorization or API tokens) before the request is sent by the
// transport client.
//
// The Apply method is invoked during request execution and should:
//
//   - mutate the provided request as needed
//   - avoid performing heavy operations
//   - avoid mutating shared state
//
// Implementations should be safe for reuse across multiple requests.
//
// Example:
//
//	client := api.New(
//	    api.WithAuth(auth.NewAccessToken("X-Access-Token", token)),
//	)
//
// The auth package intentionally defines only this minimal interface
// so authentication strategies remain transport-agnostic.
type Auth interface {
	// Apply modifies the provided HTTP request to include
	// authentication data.
	//
	// The request will be executed immediately after Apply returns,
	// so implementations should avoid long-running operations.
	//
	// Returning an error prevents the request from being executed.
	Apply(ctx context.Context, req *http.Request) error
}
