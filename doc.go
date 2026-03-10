// Package transport provides minimal primitives for executing
// requests across different communication transports.
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
// Credentials may also be resolved dynamically using
// authentication providers. Providers implement the
// `auth.Provider` interface and resolve credentials
// from configuration before a request is executed.
//
// Some providers maintain internal credential state
// (such as OAuth2 access tokens) and may optionally
// implement the `auth.Refreshable` interface to allow
// forced credential renewal.
//
// The project is designed to remain:
//
//   - small
//   - composable
//   - transport-focused
//
// allowing it to be embedded easily into larger systems.
package transport
