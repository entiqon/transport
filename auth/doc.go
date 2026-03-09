// Package auth defines the credential contract used by transport clients.
//
// A credential is responsible for applying authentication data to an
// outgoing HTTP request before it is executed by a transport.
//
// Transport implementations depend only on the Credential interface,
// allowing different authentication strategies (tokens, API keys,
// request signing, etc.) to be plugged in without coupling the
// transport layer to a specific mechanism.
package auth
