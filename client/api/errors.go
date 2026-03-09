package api

import "errors"

var (
	// InvalidRequestError indicates that the provided Request is nil
	// or otherwise invalid for execution.
	InvalidRequestError = errors.New("invalid request")

	// MissingMethodError indicates that the request method was not provided.
	MissingMethodError = errors.New("missing request method")

	// MissingPathError indicates that the request path or endpoint is missing.
	MissingPathError = errors.New("missing request path")
)
