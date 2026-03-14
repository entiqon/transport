package transport_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/entiqon/transport"
	"github.com/entiqon/transport/client/api"
	"github.com/entiqon/transport/helpers"
)

func TestBody(t *testing.T) {

	t.Run("JSON", func(t *testing.T) {
		body := transport.JSON(map[string]string{
			"name": "john",
		})

		if body == nil {
			t.Fatal("expected body")
		}

		r, err := body.Reader()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		data, err := io.ReadAll(r)
		if err != nil {
			t.Fatalf("failed reading body: %v", err)
		}

		expected := `{"name":"john"}`

		if string(data) != expected {
			t.Fatalf("expected %s, got %s", expected, string(data))
		}

		if body.ContentType() != "application/json" {
			t.Fatalf(
				"expected content type application/json, got %s",
				body.ContentType(),
			)
		}
	})

	t.Run("NewReaderEachCall", func(t *testing.T) {
		body := transport.JSON(map[string]string{
			"name": "john",
		})

		r1, err := body.Reader()
		if err != nil {
			t.Fatal(err)
		}

		r2, err := body.Reader()
		if err != nil {
			t.Fatal(err)
		}

		b1, _ := io.ReadAll(r1)
		b2, _ := io.ReadAll(r2)

		if string(b1) != string(b2) {
			t.Fatalf("reader results differ")
		}
	})

	t.Run("MarshalError", func(t *testing.T) {
		// channels cannot be marshalled to JSON
		body := transport.JSON(make(chan int))

		_, err := body.Reader()
		if err == nil {
			t.Fatal("expected marshal error")
		}
	})

	t.Run("ContentType", func(t *testing.T) {
		t.Run("Auto", func(t *testing.T) {
			server := helpers.NewTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("Content-Type") != "application/json" {
					t.Fatalf(
						"expected Content-Type application/json, got %s",
						r.Header.Get("Content-Type"),
					)
				}

				w.WriteHeader(http.StatusOK)
			})

			client := api.New()

			req := &transport.Request{
				Method: http.MethodPost,
				Path:   server.URL,
				Body: transport.JSON(map[string]string{
					"name": "john",
				}),
			}

			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}

			if !resp.OK() {
				t.Fatal("unexpected status")
			}
		})

		t.Run("Override", func(t *testing.T) {
			server := helpers.NewTestServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("Content-Type") != "application/custom" {
					t.Fatalf("header override not respected")
				}

				w.WriteHeader(http.StatusOK)
			})

			client := api.New()

			req := &transport.Request{
				Method: http.MethodPost,
				Path:   server.URL,
				Headers: map[string]string{
					"Content-Type": "application/custom",
				},
				Body: transport.JSON(map[string]string{
					"name": "john",
				}),
			}

			_, err := client.Execute(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}
		})
	})
}
