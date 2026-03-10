// Package transport provides minimal primitives for executing
// communication across different transport channels.
//
// The library focuses strictly on the communication layer,
// allowing applications to interact with external systems
// through a unified transport abstraction.
//
// transport intentionally avoids application concerns such as:
//
//   - business workflows
//   - orchestration logic
//   - domain transformations
//
// These responsibilities belong to the consuming application.
//
// Authentication is handled through pluggable credential
// strategies that modify outgoing transport requests before
// execution.
//
// Credential strategies implement the `auth.Credential`
// interface and are applied through client configuration
// using `WithCredential`.
//
// Supported credential strategies include:
//
//   - credential.AccessToken
//   - credential.APIKey
//   - credential.Basic
//   - credential.BearerToken
//   - credential.JWT
//   - credential.HMAC
//
// These strategies remain independent from the transport
// client so applications can implement their own credential
// resolution mechanisms (OAuth2 flows, token refresh,
// credential caching, etc.).
//
// The project is designed to remain:
//
//   - small
//   - composable
//   - transport-focused
//
// allowing it to be embedded easily into larger systems.
package transport
