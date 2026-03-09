package token_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/entiqon/transport/token"
)

func TestBearerToken(t *testing.T) {

	t.Run("Success", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "https://example.com", nil)

		cred := token.NewBearerToken("token")

		err := cred.Apply(context.Background(), req)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if req.Header.Get("Authorization") != "Bearer token" {
			t.Fatalf(
				"expected 'Authorization: Bearer token' but got %s",
				req.Header.Get("Authorization"),
			)
		}
	})

	t.Run("Error", func(t *testing.T) {

		req, _ := http.NewRequest("GET", "https://example.com", nil)

		cred := token.NewBearerToken("")

		err := cred.Apply(context.Background(), req)
		if err == nil {
			t.Fatal("expected error")
		}

		if req.Header.Get("Authorization") != "" {
			t.Fatalf("expected Authorization header to be empty")
		}
	})
}
