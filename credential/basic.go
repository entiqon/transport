package credential

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"

	"github.com/entiqon/transport/auth"
)

// basic implements HTTP Basic Authentication.
//
// The credential injects an Authorization header
// using the following format:
//
//	Authorization: Basic <base64(username:password)>
//
// Example:
//
//	client := api.New(
//	    api.WithCredential(
//	        token.Basic("user", "pass"),
//	    ),
//	)
type basic struct {
	username string
	password string
}

// Basic creates a BasicAuth credential.
//
// username is the account username.
// password is the account password.
func Basic(username, password string) auth.Credential {
	return &basic{
		username: username,
		password: password,
	}
}

// Apply adds the Basic Authorization header to the request.
func (b *basic) Apply(_ context.Context, req *http.Request) error {

	if b.username == "" {
		return fmt.Errorf("token: missing username")
	}

	if b.password == "" {
		return fmt.Errorf("token: missing password")
	}

	credentials := b.username + ":" + b.password
	encoded := base64.StdEncoding.EncodeToString([]byte(credentials))

	req.Header.Set("Authorization", "Basic "+encoded)

	return nil
}
