package credential

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entiqon/transport/auth"
)

// APIKeyLocation defines where an API key should be injected
// into an outgoing HTTP request.
//
// Some APIs expect the API key to be provided in a request header,
// while others require it as a query parameter. APIKeyLocation
// allows the credential to support both mechanisms.
type APIKeyLocation string

const (
	// APIKeyHeader indicates that the API key should be added
	// to the request headers.
	APIKeyHeader APIKeyLocation = "header"

	// APIKeyQuery indicates that the API key should be added
	// as a query parameter in the request URL.
	APIKeyQuery APIKeyLocation = "query"
)

// IsValid reports whether the APIKeyLocation represents a supported
// injection location.
//
// Valid locations are APIKeyHeader and APIKeyQuery.
func (l APIKeyLocation) IsValid() bool {
	switch l {
	case APIKeyHeader, APIKeyQuery:
		return true
	default:
		return false
	}
}

// apiKey implements Auth using a static API key.
//
// The API key can be injected either into an HTTP header
// or as a query parameter depending on the API requirements.
//
// Examples:
//
// Header-based authentication:
//
//	X-API-Key: <key>
//
// Authorization header:
//
//	Authorization: ApiKey <key>
//
// Query parameter:
//
//	GET /resource?api_key=<key>
type apiKey struct {
	key   string
	value string
	in    APIKeyLocation // header | query
}

// APIKey creates a new API key credential.
//
// key is the header or query parameter name (e.g. "X-API-Key").
// value is the API key value.
//
// The optional location specifies where the key should be injected.
// If omitted, the key is added to the request headers.
func APIKey(key, value string, in ...APIKeyLocation) auth.Credential {

	location := APIKeyHeader

	if len(in) > 0 {
		location = in[0]
	}

	return &apiKey{
		key:   key,
		value: value,
		in:    location,
	}
}

// Apply adds the API key to the outgoing HTTP request.
func (t *apiKey) Apply(_ context.Context, req *http.Request) error {
	if t.key == "" {
		return fmt.Errorf("token: missing API key name")
	}
	if t.value == "" {
		return fmt.Errorf("token: missing API key value")
	}

	if !t.in.IsValid() {
		return fmt.Errorf("token: invalid API key location")
	}

	switch t.in {

	case APIKeyHeader:
		req.Header.Set(t.key, t.value)

	case APIKeyQuery:
		q := req.URL.Query()
		q.Set(t.key, t.value)
		req.URL.RawQuery = q.Encode()

	}

	return nil
}
