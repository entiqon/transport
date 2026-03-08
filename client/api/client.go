package api

import (
	"context"
)

// Client defines a transport capable of executing
// communication requests to external systems.
type Client interface {
	Execute(ctx context.Context, req *Request) (*Response, error)
}
