package api_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/entiqon/transport"
	"github.com/entiqon/transport/client/api"
	"github.com/entiqon/transport/config"
	"github.com/entiqon/transport/credential"
	"github.com/entiqon/transport/provider"
)

type fakeCredential struct {
	err error
}

func (f fakeCredential) Apply(_ context.Context, _ *http.Request) error {
	return f.err
}

type failingRoundTripper struct{}

func (f failingRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("http error")
}

type failingReadCloser struct{}

func (f failingReadCloser) Read(_ []byte) (int, error) {
	return 0, errors.New("read error")
}

func (f failingReadCloser) Close() error {
	return nil
}

type failingBodyRoundTripper struct{}

func (f failingBodyRoundTripper) RoundTrip(*http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       failingReadCloser{},
		Header:     make(http.Header),
	}, nil
}

func TestAPIClient(t *testing.T) {

	t.Run("New", func(t *testing.T) {

		t.Run("Success", func(t *testing.T) {
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				if r.Header.Get("Authorization") != "Bearer test-token" {
					t.Fatalf("missing credential")
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{"ok":true}`))

			}))
			defer srv.Close()

			client := api.New(
				api.WithCredential(credential.BearerToken("test-token")),
			)

			req := &transport.Request{
				Method: http.MethodGet,
				Path:   srv.URL,
			}

			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}

			if !resp.OK() {
				t.Fatal("expected OK response")
			}
		})

		t.Run("ClientFallback", func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := api.New(api.WithHTTPClient(nil))

			req := &transport.Request{
				Method: "GET",
				Path:   server.URL,
			}

			resp, err := client.Execute(context.Background(), req)
			if err != nil {
				t.Fatal(err)
			}

			if !resp.OK() {
				t.Fatal("unexpected status")
			}
		})

		t.Run("With", func(t *testing.T) {

			t.Run("Credential", func(t *testing.T) {

				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				}))
				defer server.Close()

				client := api.New(
					api.WithCredential(fakeCredential{err: nil}),
				)

				req := &transport.Request{
					Method: "GET",
					Path:   server.URL,
				}

				_, err := client.Execute(context.Background(), req)
				if err != nil {
					t.Fatal(err)
				}
			})

			t.Run("Provider", func(t *testing.T) {

				srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
				}))
				defer srv.Close()

				oauth2 := provider.OAuth2(srv.Client())

				authConfig := config.Auth{
					Kind: "oauth2",
					OAuth2: &config.OAuth2{
						GrantType:    "client_credentials",
						TokenURL:     srv.URL,
						ClientID:     "client",
						ClientSecret: "secret",
					},
				}

				client := api.New(
					api.WithAuthProvider(oauth2, authConfig),
				)

				req := &transport.Request{
					Method: http.MethodGet,
					Path:   srv.URL,
				}

				_, _ = client.Execute(context.Background(), req)
			})
		})

		t.Run("InvalidConfig", func(t *testing.T) {

			client := api.New(
				api.WithAuthProvider(provider.OAuth2(nil), config.Auth{}), // Kind intentionally empty
			)

			req := &transport.Request{
				Method: http.MethodGet,
				Path:   "http://example.com",
			}

			_, err := client.Execute(context.Background(), req)
			if err == nil {
				t.Fatal("expected invalid auth config error")
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

			req := &transport.Request{
				Path: "https://example.com",
			}

			_, err := client.Execute(context.Background(), req)

			if err == nil {
				t.Fatal("expected error for missing method")
			}
		})

		t.Run("MissingPathError", func(t *testing.T) {

			client := api.New()

			req := &transport.Request{
				Method: "GET",
			}

			_, err := client.Execute(context.Background(), req)

			if err == nil {
				t.Fatal("expected error for missing path")
			}
		})

		t.Run("ContextCanceled", func(t *testing.T) {

			ctx, cancel := context.WithCancel(context.Background())
			cancel()

			client := api.New()

			req := &transport.Request{
				Method: "GET",
				Path:   "https://example.com",
			}

			_, err := client.Execute(ctx, req)

			if err == nil {
				t.Fatal("expected context error")
			}

			if !errors.Is(err, context.Canceled) {
				t.Fatalf("expected context.Canceled got %v", err)
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

			req := &transport.Request{
				Method: "GET",
				Path:   server.URL,
			}

			resp, err := client.Execute(context.Background(), req)

			if err != nil {
				t.Fatal(err)
			}

			if !resp.OK() {
				t.Fatalf("expected 200 got %d", resp.Status)
			}

			if resp.StatusText() != "OK" {
				t.Fatalf("unexpected status text: %s", resp.StatusText())
			}
		})

		t.Run("QueryPropagation", func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				if r.URL.Query().Get("q") != "transport" {
					t.Fatalf(
						"expected query 'transport', got '%s'",
						r.URL.Query().Get("q"),
					)
				}

				if r.URL.Query().Get("page") != "1" {
					t.Fatalf(
						"expected query '1', got '%s'",
						r.URL.Query().Get("page"),
					)
				}

				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := api.New()

			req := &transport.Request{
				Method: "GET",
				Path:   server.URL,
				Query: map[string]string{
					"q":    "transport",
					"page": "1",
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

		t.Run("HeaderPropagation", func(t *testing.T) {

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				if r.Header.Get("X-Test-Header") != "transport-test" {
					t.Fatalf(
						"expected header 'transport-test', got '%s'",
						r.Header.Get("X-Test-Header"),
					)
				}

				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := api.New()

			req := &transport.Request{
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

		t.Run("Errors", func(t *testing.T) {

			t.Run("Request", func(t *testing.T) {

				client := api.New()

				req := &transport.Request{
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

				req := &transport.Request{
					Method: "GET",
					Path:   "https://example.com",
				}

				_, err := client.Execute(context.Background(), req)

				if err == nil {
					t.Fatal("expected http error")
				}
			})

			t.Run("BodyRead", func(t *testing.T) {

				httpClient := &http.Client{
					Transport: failingBodyRoundTripper{},
				}

				client := api.New(
					api.WithHTTPClient(httpClient),
				)

				req := &transport.Request{
					Method: "GET",
					Path:   "https://example.com",
				}

				_, err := client.Execute(context.Background(), req)

				if err == nil {
					t.Fatal("expected body read error")
				}
			})
		})
	})

	t.Run("Error", func(t *testing.T) {

		t.Run("Credential", func(t *testing.T) {

			client := api.New(
				api.WithCredential(fakeCredential{err: errors.New("credential failed")}),
			)

			req := &transport.Request{
				Method: "GET",
				Path:   "https://example.com",
			}

			_, err := client.Execute(context.Background(), req)

			if err == nil {
				t.Fatal("expected credential error")
			}
		})
	})
}
