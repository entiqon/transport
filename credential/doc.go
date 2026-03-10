// Package credential provides authentication credential strategies
// used by the transport library.
//
// Credential strategies implement the auth.Credential interface and
// inject authentication data into outgoing HTTP requests.
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
//   - Basic
//     Injects the Authorization header using the
//     HTTP Basic authentication scheme.
//
//   - JWT
//     Injects a JSON Web Token (JWT) into an outgoing HTTP request.
//     If the header is "Authorization", the Bearer scheme is applied
//     automatically.
//
//   - HMAC
//     Signs outgoing HTTP requests using an HMAC-SHA256 signature.
//     The credential injects authentication headers including:
//
//     X-Key
//     X-Timestamp
//     X-Signature
//
//     The signature is computed from request metadata using a
//     shared secret.
//
// These strategies remain independent of the transport client,
// allowing authentication mechanisms to evolve without modifying
// the underlying HTTP transport implementation.
package credential
