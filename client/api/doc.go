// Package api provides an HTTP transport client implementation for
// the transport library.
//
// The API client is responsible for executing HTTP requests while
// remaining independent of authentication mechanisms and application
// business logic.
//
// Authentication is applied through strategies implementing the
// auth.Credential interface. These strategies mutate outgoing HTTP
// requests before execution.
//
// Credentials may be provided directly or resolved dynamically using
// authentication providers.
//
// Providers implement the auth.Provider interface and resolve
// credentials from configuration before a request is executed.
// Some providers may implement auth.Refreshable to support credential
// renewal when tokens expire.
//
// The client supports versioned APIs through the WithVersion option.
// When configured, the version segment is inserted between the base
// URL and request path.
//
// Examples:
//
// Using a Bearer token:
//
//	client := api.New(
//	    api.WithCredential(
//	        credential.BearerToken("token"),
//	    ),
//	)
//
// Using an OAuth2 provider:
//
//	client := api.New(
//	    api.WithAuthProvider(
//	        provider.OAuth2(http.DefaultClient),
//	        authConfig,
//	    ),
//	)
//
// The client validates transport requests, constructs an http.Request,
// applies credentials if configured, executes the request using the
// configured http.Client, and returns a transport.Response.
package api
