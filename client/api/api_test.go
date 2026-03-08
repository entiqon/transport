package api_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/entiqon/transport/client/api"
)

type successAuth struct{}

func (successAuth) Apply(_ context.Context, r *http.Request) error {
	r.Header.Set("Authorization", "Bearer test")
	return nil
}

type fakeAuth struct {
	err error
}

func (f fakeAuth) Apply(_ context.Context, _ *http.Request) error {
	return f.err
}

type failingRoundTripper struct{}

func (f failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("http error")
}

func TestAPIClient(t *testing.T) {

	t.Run("New", func(t *testing.T) {

		t.Run("ClientFallback", func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := api.New(api.WithHTTPClient(nil))

			req := &api.Request{
				Method: "GET",
				Path:   server.URL,
			}

			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}

			if resp.Status != http.StatusOK {
				t.Fatal("unexpected status")
			}
		})
	})

	t.Run("Validation", func(t *testing.T) {

		t.Run("InvalidRequestError", func(t *testing.T) {

			client := api.New()

			_, err := client.Execute(context.Background(), nil)

			if err == nil {
				t.Fatal("expected error for nil request")
			}
		})

		t.Run("MissingMethodError", func(t *testing.T) {

			client := api.New()

			req := &api.Request{
				Path: "http://example.com",
			}

			_, err := client.Execute(context.Background(), req)

			if err == nil {
				t.Fatal("expected error for missing method")
			}
		})

		t.Run("MissingPathError", func(t *testing.T) {

			client := api.New()

			req := &api.Request{
				Method: "GET",
			}

			_, err := client.Execute(context.Background(), req)

			if err == nil {
				t.Fatal("expected error for missing path")
			}
		})
	})

	t.Run("Execution", func(t *testing.T) {

		t.Run("Success", func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := api.New()

			req := &api.Request{
				Method: "GET",
				Path:   server.URL,
			}

			resp, err := client.Execute(context.Background(), req)

			if err != nil {
				t.Fatal(err)
			}

			if resp.Status != http.StatusOK {
				t.Fatalf("expected 200 got %d", resp.Status)
			}
		})

		t.Run("Headers", func(t *testing.T) {

			t.Run("Propagation", func(t *testing.T) {

				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

					if r.Header.Get("X-Test-Header") != "transport-test" {
						t.Fatalf("expected header 'transport-test', got '%s'",
							r.Header.Get("X-Test-Header"))
					}

					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()

				client := api.New()

				req := &api.Request{
					Method: "GET",
					Path:   server.URL,
					Headers: map[string]string{
						"X-Test-Header": "transport-test",
					},
				}

				resp, err := client.Execute(context.Background(), req)
				if err != nil {
					t.Fatal(err)
				}

				if resp.Status != http.StatusOK {
					t.Fatalf("unexpected status: %d", resp.Status)
				}
			})
		})

		t.Run("Auth", func(t *testing.T) {

			t.Run("Success", func(t *testing.T) {

				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

					if r.Header.Get("Authorization") != "Bearer test" {
						t.Fatal("authorization header missing")
					}

					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()

				client := api.New(api.WithAuth(successAuth{}))

				req := &api.Request{
					Method: "GET",
					Path:   server.URL,
				}

				_, err := client.Execute(context.Background(), req)
				if err != nil {
					t.Fatal(err)
				}
			})

			t.Run("Error", func(t *testing.T) {

				client := api.New(
					api.WithAuth(fakeAuth{err: errors.New("auth failed")}),
				)

				req := &api.Request{
					Method: "GET",
					Path:   "http://example.com",
				}

				_, err := client.Execute(context.Background(), req)

				if err == nil {
					t.Fatal("expected auth error")
				}
			})
		})

		t.Run("Errors", func(t *testing.T) {
			t.Run("Request", func(t *testing.T) {

				client := api.New()

				req := &api.Request{
					Method: "GET",
					Path:   "://bad-url",
				}

				_, err := client.Execute(context.Background(), req)

				if err == nil {
					t.Fatal("expected url parsing error")
				}
			})

			t.Run("HTTP", func(t *testing.T) {

				httpClient := &http.Client{
					Transport: failingRoundTripper{},
				}

				client := api.New(
					api.WithHTTPClient(httpClient),
				)

				req := &api.Request{
					Method: "GET",
					Path:   "http://example.com",
				}

				_, err := client.Execute(context.Background(), req)

				if err == nil {
					t.Fatal("expected http error")
				}
			})
		})
	})
}
