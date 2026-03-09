package api

import (
	"context"
)

// Client represents a transport capable of executing
// communication requests to external systems.
//
// Implementations are responsible for translating a Request
// into a concrete protocol operation (such as an HTTP request),
// executing it, and returning the resulting Response.
type Client interface {

	// Execute performs the provided transport Request.
	//
	// The request is validated, transformed into the underlying
	// protocol representation, and executed using the configured
	// transport mechanism.
	//
	// If the request cannot be executed, an error is returned.
	Execute(ctx context.Context, req *Request) (*Response, error)
}
