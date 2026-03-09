package api

import "io"

// Request represents a transport request executed by a Client.
//
// It describes the information required to perform an outbound
// communication request independently of a specific transport
// implementation. The client is responsible for translating this
// structure into the underlying protocol request (e.g., HTTP).
type Request struct {

	// Connection identifies the logical connection or integration
	// context associated with the request. Higher-level components
	// may use this value to resolve endpoints or credentials.
	Connection string

	// Method defines the request method (e.g., GET, POST, PUT, DELETE).
	Method string

	// Path defines the target resource path or full endpoint URL.
	Path string

	// Headers contains optional HTTP headers applied to the request.
	Headers map[string]string

	// Query contains optional query parameters appended to the request URL.
	Query map[string]string

	// Body represents the request payload.
	//
	// The value must implement io.Reader and will be passed directly
	// to the underlying transport request.
	Body io.Reader
}
