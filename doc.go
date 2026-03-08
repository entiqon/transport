// Package transport provides minimal primitives for executing
// communication across different transport channels.
//
// The library focuses strictly on the communication layer,
// allowing applications to interact with external systems
// through a unified transport abstraction.
//
// transport intentionally avoids application concerns such as:
//   - business workflows
//   - orchestration logic
//   - domain transformations
//
// These responsibilities belong to the consuming application.
//
// The project is designed to remain small, composable,
// and transport-focused.
package transport
