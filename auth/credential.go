package auth

import (
	"context"
	"net/http"
)

// Credential represents a strategy that applies authentication
// data to an outgoing HTTP request.
//
// Implementations typically mutate request headers (for example,
// Authorization or API tokens) before the request is executed by
// the transport client.
//
// The Apply method is invoked during request execution and should:
//
//   - mutate the provided request as needed
//   - avoid blocking or network operations
//   - avoid mutating shared state
//
// Implementations should be safe for reuse across multiple requests.
//
// Example:
//
//	client := api.New(
//	    api.WithCredential(token.NewAccessToken("X-Access-Token", token)),
//	)
//
// The auth package intentionally defines only this minimal interface
// so credential strategies remain transport-agnostic.
type Credential interface {

	// Apply modifies the provided HTTP request to include
	// authentication data.
	//
	// The request will be executed immediately after Apply returns,
	// so implementations should avoid long-running or blocking work.
	//
	// Returning an error prevents the request from being executed.
	Apply(ctx context.Context, req *http.Request) error
}
