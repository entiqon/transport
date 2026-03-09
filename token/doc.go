// Package token provides token-based authentication strategies
// used by the transport library.
//
// Token strategies implement the auth.Credential interface and inject
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
//   - APIKey
//     Injects an API key either as an HTTP header or
//     as a query parameter.
//
// These strategies remain independent of the transport client,
// allowing authentication mechanisms to evolve without modifying
// the underlying HTTP transport implementation.
package token
