package credential_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/entiqon/transport/credential"
)

func TestAPIKeyToken(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		t.Run("Header", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			cred := credential.APIKey("X-API-Key", "token", credential.APIKeyHeader)

			err := cred.Apply(context.Background(), req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if req.Header.Get("X-API-Key") != "token" {
				t.Fatalf(
					"expected 'X-API-Key: token' but got %s",
					req.Header.Get("X-API-Key"),
				)
			}
		})

		t.Run("DefaultHeader", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			// location omitted -> should default to header
			cred := credential.APIKey("X-API-Key", "token")

			err := cred.Apply(context.Background(), req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if req.Header.Get("X-API-Key") != "token" {
				t.Fatalf(
					"expected default header 'X-API-Key: token' but got %s",
					req.Header.Get("X-API-Key"),
				)
			}
		})

		t.Run("QueryParams", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			cred := credential.APIKey("api_key", "token", credential.APIKeyQuery)

			err := cred.Apply(context.Background(), req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if req.URL.Query().Get("api_key") != "token" {
				t.Fatalf(
					"expected query parameter api_key=token but got %s",
					req.URL.Query().Get("api_key"),
				)
			}
		})
	})

	t.Run("Error", func(t *testing.T) {

		t.Run("MissingKey", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			cred := credential.APIKey("", "token")

			err := cred.Apply(context.Background(), req)
			if err == nil {
				t.Fatal("expected error for missing API key name")
			}
		})

		t.Run("MissingValue", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			cred := credential.APIKey("X-API-Key", "")

			err := cred.Apply(context.Background(), req)
			if err == nil {
				t.Fatal("expected error for missing API key value")
			}
		})

		t.Run("InvalidLocation", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			cred := credential.APIKey(
				"X-API-Key",
				"token",
				credential.APIKeyLocation("invalid"),
			)

			err := cred.Apply(context.Background(), req)
			if err == nil {
				t.Fatal("expected error for invalid API key location")
			}
		})
	})
}
