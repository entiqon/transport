package transport

import "io"

// Body represents a serializable request payload.
//
// Implementations provide the request reader and declare the associated
// Content-Type. The transport client uses this information when building
// the underlying http.Request.
//
// Implementations should:
//
//   - Return a new reader on every Reader() call
//   - Avoid mutating internal state
//   - Return a non-empty Content-Type when applicable
//
// Example:
//
//	req := &transport.Request{
//	    Method: http.MethodPost,
//	    Path:   "/users",
//	    Body:   transport.JSON(user),
//	}
type Body interface {

	// Reader returns the serialized request payload.
	Reader() (io.Reader, error)

	// ContentType returns the HTTP Content-Type associated with the body.
	//
	// The transport client uses this value to automatically populate the
	// Content-Type header when executing requests.
	//
	// If the request already defines a Content-Type header, that value
	// takes precedence over the body-provided content type.
	ContentType() string
}
