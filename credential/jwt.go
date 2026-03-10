package credential

import (
	"context"
	"errors"
	"net/http"

	"github.com/entiqon/transport/auth"
)

// JWT returns a credential that injects a JSON Web Token (JWT)
// into an outgoing HTTP request.
//
// The token is written to the specified header. If the header is
// "Authorization", the credential automatically applies the
// Bearer authentication scheme:
//
//	Authorization: Bearer <token>
//
// For other headers, the token is written as-is:
//
//	X-JWT-Assertion: <token>
//
// JWT only injects the token and does not generate, sign, validate,
// or refresh it. Those responsibilities belong to higher-level
// credential resolvers in the consuming application.
func JWT(header, token string) auth.Credential {
	return &jwt{
		header: header,
		token:  token,
	}
}

// jwt implements JWT-based authentication for outgoing requests.
type jwt struct {
	header string
	token  string
}

// Apply injects the JWT into the HTTP request.
//
// Validation rules:
//
//   - header must not be empty
//   - token must not be empty
//
// If the header is "Authorization", the Bearer scheme is applied.
// Otherwise, the token is written directly to the configured header.
func (j *jwt) Apply(_ context.Context, req *http.Request) error {

	if j.header == "" {
		return errors.New("credential: JWT header cannot be empty")
	}

	if j.token == "" {
		return errors.New("credential: JWT token cannot be empty")
	}

	if j.header == "Authorization" {
		req.Header.Set(j.header, "Bearer "+j.token)
		return nil
	}

	req.Header.Set(j.header, j.token)

	return nil
}
