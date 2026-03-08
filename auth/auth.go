package auth

import "net/http"

// Auth defines a mechanism that can apply authentication
// to an outgoing HTTP request.
type Auth interface {
	Apply(req *http.Request) error
}
