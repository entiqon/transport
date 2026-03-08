package api

import "io"

// Request represents a transport request executed by a Client.
//
// It describes the essential information required to perform an HTTP
// request independently of a specific transport implementation.
// The client is responsible for translating this structure into the
// underlying protocol request (e.g., HTTP).
type Request struct {

	// Connection identifies the logical connection or integration
	// configuration used to resolve the target endpoint.
	Connection string

	// Method defines the request method (e.g., GET, POST, PUT, DELETE).
	Method string

	// Path defines the target resource path or full endpoint URL.
	Path string

	// Headers contains optional request headers applied to the request.
	Headers map[string]string

	// Query contains optional query parameters appended to the request URL.
	Query map[string]string

	// Body represents the request payload.
	// The concrete type depends on the transport client implementation
	// and may be encoded before sending the request.
	Body io.Reader
}
