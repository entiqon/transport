package credential_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/entiqon/transport/credential"
)

func TestAccessToken(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "https://example.com", nil)

		cred := credential.AccessToken("X-Access-Token", "token")

		err := cred.Apply(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if req.Header.Get("X-Access-Token") != "token" {
			t.Fatalf("expected header to be set")
		}
	})

	t.Run("Error", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "https://example.com", nil)

		cred := credential.AccessToken("", "token")

		err := cred.Apply(context.Background(), req)
		if err == nil || err.Error() != "api key missing" {
			t.Fatal("expected error")
		}

		if req.Header.Get("X-Access-Token") != "" {
			t.Fatal("expected header to be empty")
		}
	})
}
