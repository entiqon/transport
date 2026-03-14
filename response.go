package transport

import (
	"encoding/json"
	"net/http"
)

// Response represents the result of executing a transport Request.
//
// It contains normalized response information together with the raw
// protocol response returned by the underlying transport implementation.
type Response struct {

	// Status is the HTTP status code returned by the server.
	Status int

	// Headers contains the normalized response headers returned by the server.
	//
	// Multi-value headers may be flattened depending on the transport
	// implementation.
	Headers map[string]string

	// Body contains the raw response payload returned by the server.
	Body []byte

	// Raw contains the original HTTP response returned by the underlying
	// transport implementation.
	//
	// The response body is consumed by the transport client when building
	// the Response structure, therefore Raw.Body should not be read.
	// The normalized Body field contains the full payload.
	Raw *http.Response
}

// OK reports whether the response status code indicates success
// according to HTTP semantics (status code in the 2xx range).
func (r *Response) OK() bool {
	return r.Status >= 200 && r.Status < 300
}

// Header returns the value of the specified response header.
//
// If the header is not present, an empty string is returned.
func (r *Response) Header(name string) string {
	if r.Headers == nil {
		return ""
	}
	return r.Headers[name]
}

// StatusText returns the standard HTTP status text associated with
// the response status code.
//
// It is equivalent to calling http.StatusText on the Status field
// and is provided as a convenience helper when logging or reporting
// transport responses.
func (r *Response) StatusText() string {
	return http.StatusText(r.Status)
}

// JSON decodes the response body as JSON into the provided value.
//
// It is a convenience helper that avoids manually calling json.Unmarshal
// when working with JSON APIs.
//
// Example:
//
//	var order Order
//	if err := resp.JSON(&order); err != nil {
//	    return err
//	}
func (r *Response) JSON(out any) error {
	return json.Unmarshal(r.Body, out)
}
