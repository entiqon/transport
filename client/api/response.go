package api

import "net/http"

// Response represents the result of executing a transport Request.
//
// It contains the HTTP status code, normalized response headers,
// and the raw response payload returned by the underlying transport
// implementation.
type Response struct {
	// Status is the HTTP status code returned by the server.
	Status int

	// Headers contains the response headers returned by the server.
	Headers map[string]string

	// Body contains the raw response payload returned by the server.
	Body []byte

	// Raw contains the original HTTP response returned by the underlying
	// transport implementation. It may be nil if the transport does not
	// expose the underlying protocol response.
	//
	// This field is intended for advanced use cases such as debugging,
	// inspecting protocol-specific details, or accessing features not
	// represented in the normalized Response fields.
	Raw *http.Response
}

// OK reports whether the response status code indicates success.
func (r *Response) OK() bool {
	return r.Status >= 200 && r.Status < 300
}
