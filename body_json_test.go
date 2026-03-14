package transport_test

import (
	"io"
	"testing"

	"github.com/entiqon/transport"
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
}
