package transport_test

import (
	"testing"

	"github.com/entiqon/transport"
)

func TestResponse(t *testing.T) {
	t.Run("Header", func(t *testing.T) {
		t.Run("Nil", func(t *testing.T) {
			resp := &transport.Response{}

			if resp.Header("Content-Type") != "" {
				t.Fatal("expected empty header for nil headers map")
			}
		})

		t.Run("Value", func(t *testing.T) {
			resp := &transport.Response{
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
	})

	t.Run("Status", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			resp := &transport.Response{Status: 200}

			if !resp.OK() {
				t.Fatal("expected OK for status 200")
			}
		})

		t.Run("Text", func(t *testing.T) {
			resp := &transport.Response{Status: 200}

			if resp.StatusText() != "OK" {
				t.Fatalf("unexpected status text: %s", resp.StatusText())
			}
		})
	})

	t.Run("Body", func(t *testing.T) {
		t.Run("Nil", func(t *testing.T) {
			resp := &transport.Response{}

			var out map[string]any

			if err := resp.JSON(&out); err == nil {
				t.Fatal("expected JSON error for empty body")
			}
		})
	})

	t.Run("JSON", func(t *testing.T) {
		t.Run("Decode", func(t *testing.T) {
			resp := &transport.Response{
				Body: []byte(`{"name":"john"}`),
			}

			var out struct {
				Name string `json:"name"`
			}

			if err := resp.JSON(&out); err != nil {
				t.Fatalf("unexpected error decoding JSON: %v", err)
			}

			if out.Name != "john" {
				t.Fatalf("expected name 'john', got '%s'", out.Name)
			}
		})

		t.Run("Invalid", func(t *testing.T) {
			resp := &transport.Response{
				Body: []byte(`{invalid json}`),
			}

			var out map[string]any

			if err := resp.JSON(&out); err == nil {
				t.Fatal("expected JSON decoding error")
			}
		})
	})
}
