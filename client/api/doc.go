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
// Examples:
//
// Using an Access Token header:
//
//	client := api.New(
//	    api.WithCredential(
//	        token.NewAccessToken("X-Access-Token", "token"),
//	    ),
//	)
//
// Using a Bearer token:
//
//	client := api.New(
//	    api.WithCredential(
//	        token.NewBearerToken("token"),
//	    ),
//	)
//
// Using an API key:
//
//	client := api.New(
//	    api.WithCredential(
//	        token.NewAPIKey("X-API-Key", "token", token.APIKeyHeader),
//	    ),
//	)
//
// The client validates requests, constructs an http.Request,
// applies credentials if configured, and executes the request
// using the configured http.Client.
package api
