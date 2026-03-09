// Package api provides an HTTP transport client for the transport library.
//
// The API client focuses exclusively on executing HTTP requests while
// remaining independent of authentication mechanisms and application
// business logic.
//
// Authentication data is applied through strategies implementing the
// auth.Credential interface. These strategies modify outgoing requests
// before execution.
//
// Example:
//
//	client := api.New(
//	    api.WithCredential(token.NewBearerToken("token")),
//	)
//
// The client validates requests, constructs an http.Request,
// applies credentials if configured, and executes the request
// using the configured http.Client.
package api
