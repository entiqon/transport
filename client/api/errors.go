package api

import "errors"

var (
	// InvalidRequestError indicates that the provided request is nil
	// or otherwise not valid for execution.
	InvalidRequestError = errors.New("invalid request")

	// MissingMethodError indicates that the request method was not provided.
	MissingMethodError = errors.New("missing request method")

	// MissingPathError indicates that the request path or endpoint is missing.
	MissingPathError = errors.New("missing request path")

	// InvalidBodyError indicates that the request body type is not supported
	// by the transport client. The body must implement io.Reader.
	InvalidBodyError = errors.New("invalid request body: expected io.Reader")
)
