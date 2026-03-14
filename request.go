package transport

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

	// Path defines either a relative resource path or a full endpoint URL.
	Path string

	// Headers contains optional HTTP headers applied to the request.
	Headers map[string]string

	// Query contains optional query parameters appended to the request URL.
	Query map[string]string

	// Body represents the request payload.
	//
	// The value must implement the Body interface, which allows the
	// transport client to obtain a new reader for each request execution.
	// This enables retry-safe request bodies and supports helpers such
	// as JSON payload serialization.
	//
	// If the body declares a Content-Type, the client will automatically
	// propagate it to the outgoing request unless the header is already
	// defined in Headers.
	Body Body
}
