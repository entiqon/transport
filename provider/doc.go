// Package providers contains implementations of the auth.Provider interface.
//
// Providers resolve authentication configuration into credentials that can
// be applied to outgoing HTTP requests. Each provider is responsible for
// handling the lifecycle of its authentication mechanism, such as retrieving
// tokens, refreshing expired credentials, or managing credential rotation.
//
// The transport client invokes providers during request execution to obtain
// a credential which is then applied to the HTTP request.
package provider
