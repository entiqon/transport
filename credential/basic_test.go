package credential_test

import (
	"context"
	"encoding/base64"
	"net/http"
	"testing"

	"github.com/entiqon/transport/credential"
)

func TestBasicAuthToken(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "https://example.com", nil)

		cred := credential.Basic("user", "pass")

		err := cred.Apply(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		expected := "Basic " + base64.StdEncoding.EncodeToString([]byte("user:pass"))

		if req.Header.Get("Authorization") != expected {
			t.Fatalf(
				"expected Authorization %q but got %q",
				expected,
				req.Header.Get("Authorization"),
			)
		}
	})

	t.Run("Error", func(t *testing.T) {

		t.Run("MissingUsername", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			cred := credential.Basic("", "pass")

			err := cred.Apply(context.Background(), req)

			if err == nil {
				t.Fatal("expected error for missing username")
			}
		})

		t.Run("MissingPassword", func(t *testing.T) {

			req, _ := http.NewRequest("GET", "https://example.com", nil)

			cred := credential.Basic("user", "")

			err := cred.Apply(context.Background(), req)

			if err == nil {
				t.Fatal("expected error for missing password")
			}
		})
	})
}
