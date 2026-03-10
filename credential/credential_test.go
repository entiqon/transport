package credential_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/entiqon/transport/auth"
	"github.com/entiqon/transport/credential"
)

type credentialCase struct {
	name      string
	cred      auth.Credential
	header    string
	expected  string
	expectErr bool
}

func TestCredentials(t *testing.T) {

	tests := []credentialCase{
		{
			name:     "AccessToken",
			cred:     credential.AccessToken("X-Access-Token", "token"),
			header:   "X-Access-Token",
			expected: "token",
		},
		{
			name:      "AccessTokenInvalidHeader",
			cred:      credential.AccessToken("", "token"),
			expectErr: true,
		},
		{
			name:     "BearerToken",
			cred:     credential.BearerToken("token"),
			header:   "Authorization",
			expected: "Bearer token",
		},
		{
			name:      "BearerTokenEmptyToken",
			cred:      credential.BearerToken(""),
			expectErr: true,
		},
		{
			name:     "APIKeyHeader",
			cred:     credential.APIKey("X-API-Key", "key", credential.APIKeyHeader),
			header:   "X-API-Key",
			expected: "key",
		},
		{
			name:      "APIKeyMissingKey",
			cred:      credential.APIKey("", "key", credential.APIKeyHeader),
			expectErr: true,
		},
		{
			name:      "APIKeyMissingValue",
			cred:      credential.APIKey("X-API-Key", "", credential.APIKeyHeader),
			expectErr: true,
		},
		{
			name:      "APIKeyInvalidLocation",
			cred:      credential.APIKey("X-API-Key", "key", credential.APIKeyLocation("invalid")),
			expectErr: true,
		},
		{
			name:     "BasicAuth",
			cred:     credential.Basic("user", "pass"),
			header:   "Authorization",
			expected: "Basic dXNlcjpwYXNz",
		},
		{
			name:      "BasicAuthMissingUsername",
			cred:      credential.Basic("", "pass"),
			expectErr: true,
		},
		{
			name:      "BasicAuthMissingPassword",
			cred:      credential.Basic("user", ""),
			expectErr: true,
		},
		{
			name:     "JWTAuthorization",
			cred:     credential.JWT("Authorization", "jwt"),
			header:   "Authorization",
			expected: "Bearer jwt",
		},
		{
			name:     "JWTCustomHeader",
			cred:     credential.JWT("X-JWT-Assertion", "jwt"),
			header:   "X-JWT-Assertion",
			expected: "jwt",
		},
		{
			name:      "JWTMissingHeader",
			cred:      credential.JWT("", "jwt"),
			expectErr: true,
		},
		{
			name:      "JWTMissingToken",
			cred:      credential.JWT("Authorization", ""),
			expectErr: true,
		},
	}

	for _, tc := range tests {

		t.Run(tc.name, func(t *testing.T) {

			req, err := http.NewRequest("GET", "https://example.com", nil)
			if err != nil {
				t.Fatalf("unexpected request creation error: %v", err)
			}

			err = tc.cred.Apply(context.Background(), req)

			if tc.expectErr {
				if err == nil {
					t.Fatalf("expected error")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			got := req.Header.Get(tc.header)

			if got != tc.expected {
				t.Fatalf(
					"expected header %q value %q but got %q",
					tc.header,
					tc.expected,
					got,
				)
			}
		})
	}
}
