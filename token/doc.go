// Package token provides token-based authentication strategies
// used by the transport library.
//
// Token strategies implement the auth.Auth interface and inject
// credentials into outgoing HTTP requests.
//
// Supported strategies:
//
//   - AccessToken
//     Injects a custom header containing an access token.
//
//   - BearerToken
//     Injects the standard Authorization header using the
//     Bearer authentication scheme.
//
// These strategies remain independent of the transport client,
// allowing authentication mechanisms to evolve without modifying
// the underlying HTTP transport implementation.
package token
