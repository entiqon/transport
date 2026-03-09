package api_test

import (
	"testing"

	"github.com/entiqon/transport/client/api"
)

func TestResponse(t *testing.T) {
	t.Run("HeaderNil", func(t *testing.T) {

		resp := &api.Response{}

		if resp.Header("Content-Type") != "" {
			t.Fatal("expected empty header for nil headers map")
		}
	})

	t.Run("HeaderValue", func(t *testing.T) {

		resp := &api.Response{
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		if resp.Header("Content-Type") != "application/json" {
			t.Fatalf(
				"expected header 'application/json', got '%s'",
				resp.Header("Content-Type"),
			)
		}
	})

	t.Run("OK", func(t *testing.T) {

		resp := &api.Response{Status: 200}

		if !resp.OK() {
			t.Fatal("expected OK for status 200")
		}
	})

	t.Run("StatusText", func(t *testing.T) {

		resp := &api.Response{Status: 200}

		if resp.StatusText() != "OK" {
			t.Fatalf("unexpected status text: %s", resp.StatusText())
		}
	})
}
