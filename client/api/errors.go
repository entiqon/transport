package api

import "errors"

var (
	InvalidRequestError = errors.New("api: invalid request")
	MissingMethodError  = errors.New("api: invalid request method")
	MissingPathError    = errors.New("api: invalid request path")
)
