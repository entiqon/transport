// Package auth defines the authentication contracts used by the transport layer.
//
// The package provides three core abstractions:
//
//   - Credential: modifies outgoing HTTP requests with authentication data
//   - Provider: resolves credentials dynamically from configuration
//   - Refreshable: optionally invalidates cached credentials
//
// Credential implementations mutate the HTTP request directly. Examples include:
//
//   - Bearer tokens
//   - API keys
//   - HMAC signatures
//   - JWT headers
//
// Providers resolve credentials at runtime based on configuration. This enables
// dynamic authentication flows such as OAuth2 token refresh without coupling
// transport clients to specific authentication mechanisms.
//
// Providers that maintain internal credential state (such as OAuth2 access tokens)
// may optionally implement the Refreshable interface to allow forced credential
// renewal.
//
// The transport client interacts only with these contracts, allowing different
// authentication strategies to be plugged in without modifying the client itself.
package auth
