package api_test

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/entiqon/transport"
	"github.com/entiqon/transport/client/api"
	"github.com/entiqon/transport/config"
	"github.com/entiqon/transport/credential"
	apitest2 "github.com/entiqon/transport/internal/apitest"
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

type failingBody struct{}

func (f failingBody) Reader() (io.Reader, error) { return nil, errors.New("reader error") }
func (f failingBody) ContentType() string        { return "" }

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
			server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("Authorization") != "Bearer test-token" {
					t.Fatalf("missing credential")
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(`{"ok":true}`))
			})

			client := api.New(
				api.WithCredential(credential.BearerToken("test-token")),
			)

			req := apitest2.NewRequest(http.MethodGet, server.URL)

			resp := apitest2.Execute(t, client, req)

			apitest2.AssertOK(t, resp)
		})

		t.Run("ClientFallback", func(t *testing.T) {
			server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})

			client := api.New(api.WithHTTPClient(nil))

			req := apitest2.NewRequest(http.MethodGet, server.URL)

			resp := apitest2.Execute(t, client, req)

			apitest2.AssertOK(t, resp)
		})

		t.Run("With", func(t *testing.T) {
			t.Run("Credential", func(t *testing.T) {
				server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusOK)
				})

				client := api.New(
					api.WithCredential(fakeCredential{err: nil}),
				)

				req := apitest2.NewRequest(http.MethodGet, server.URL)

				apitest2.Execute(t, client, req)
			})

			t.Run("Provider", func(t *testing.T) {
				server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
					w.Header().Set("Content-Type", "application/json")

					_, _ = w.Write([]byte(`{
						"access_token": "test-token",
						"token_type": "Bearer",
						"expires_in": 3600
					}`))
				})

				oauth2 := provider.OAuth2(server.Client())

				authConfig := config.Auth{
					Kind: "oauth2",
					OAuth2: &config.OAuth2{
						GrantType:    "client_credentials",
						TokenURL:     server.URL,
						ClientID:     "client",
						ClientSecret: "secret",
					},
				}

				client := api.New(
					api.WithAuthProvider(oauth2, authConfig),
				)

				req := apitest2.NewRequest(http.MethodGet, server.URL)

				resp := apitest2.Execute(t, client, req)

				apitest2.AssertOK(t, resp)
			})

			t.Run("BasePath", func(t *testing.T) {
				server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/api" {
						t.Fatalf(
							"expected path '/api', got '%s'",
							r.URL.Path,
						)
					}

					w.WriteHeader(http.StatusOK)
				})

				client := api.New(
					api.WithBasePath("api"),
				)

				req := apitest2.NewRequest(http.MethodGet, server.URL)

				resp := apitest2.Execute(t, client, req)

				apitest2.AssertOK(t, resp)
			})

			t.Run("Version", func(t *testing.T) {
				server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
					if r.Header.Get("X-API-Version") != "v1" {
						t.Fatalf(
							"expected version header 'v1', got '%s'",
							r.Header.Get("X-API-Version"),
						)
					}

					w.WriteHeader(http.StatusOK)
				})

				client := api.New(
					api.WithVersion("v1"),
				)

				req := apitest2.NewRequest(http.MethodGet, server.URL)

				resp := apitest2.Execute(t, client, req)

				apitest2.AssertOK(t, resp)
			})

			t.Run("BasePathAndVersion", func(t *testing.T) {
				server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/api/v1" {
						t.Fatalf("expected path '/api/v1', got '%s'", r.URL.Path)
					}

					if r.Header.Get("X-API-Version") != "v1" {
						t.Fatalf(
							"expected version header 'v1', got '%s'",
							r.Header.Get("X-API-Version"),
						)
					}

					w.WriteHeader(http.StatusOK)
				})

				client := api.New(
					api.WithBasePath("api"),
					api.WithVersion("v1"),
				)

				req := apitest2.NewRequest(http.MethodGet, server.URL)

				resp := apitest2.Execute(t, client, req)

				apitest2.AssertOK(t, resp)
			})

			t.Run("BasePathAlreadyPresent", func(t *testing.T) {
				server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
					if r.URL.Path != "/api/users" {
						t.Fatalf(
							"expected path '/api/users', got '%s'",
							r.URL.Path,
						)
					}

					w.WriteHeader(http.StatusOK)
				})

				client := api.New(
					api.WithBasePath("api"),
				)

				req := apitest2.NewRequest(http.MethodGet, server.URL+"/api/users")

				resp := apitest2.Execute(t, client, req)

				apitest2.AssertOK(t, resp)
			})
		})

		t.Run("InvalidConfig", func(t *testing.T) {
			client := api.New(
				api.WithAuthProvider(provider.OAuth2(nil), config.Auth{}), // Kind intentionally empty
			)

			req := apitest2.NewRequest(http.MethodGet, "https://example.com")

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

			req := apitest2.NewRequest(http.MethodGet, "https://example.com")

			_, err := client.Execute(ctx, req)
			if err == nil {
				t.Fatal("expected context error")
			}

			if !errors.Is(err, context.Canceled) {
				t.Fatalf("expected context.Canceled got %v", err)
			}
		})
	})

	t.Run("DoJSON", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Header.Get("Content-Type") != "application/json" {
					t.Fatalf("expected JSON content type, got %s", r.Header.Get("Content-Type"))
				}

				body, err := io.ReadAll(r.Body)
				if err != nil {
					t.Fatal(err)
				}

				if string(body) != `{"name":"john"}` {
					t.Fatalf("unexpected body: %s", string(body))
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

			if resp.Status != http.StatusOK {
				t.Fatalf("unexpected status: %d", resp.Status)
			}
		})

		t.Run("Decode", func(t *testing.T) {
			server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"name":"john"}`))
			})

			client := api.New()

			req := apitest2.NewRequest(http.MethodGet, server.URL)

			var out struct {
				Name string `json:"name"`
			}

			err := client.DoJSON(context.Background(), req, &out)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if out.Name != "john" {
				t.Fatalf("expected name 'john', got '%s'", out.Name)
			}
		})

		t.Run("InvalidJSON", func(t *testing.T) {
			server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{invalid json}`))
			})

			client := api.New()

			req := apitest2.NewRequest(http.MethodGet, server.URL)

			var out map[string]any

			err := client.DoJSON(context.Background(), req, &out)

			if err == nil {
				t.Fatal("expected JSON decode error")
			}
		})

		t.Run("Error", func(t *testing.T) {
			httpClient := &http.Client{
				Transport: failingRoundTripper{},
			}

			client := api.New(
				api.WithHTTPClient(httpClient),
			)

			req := apitest2.NewRequest(http.MethodGet, "https://example.com")

			var out map[string]any

			err := client.DoJSON(context.Background(), req, &out)

			if err == nil {
				t.Fatal("expected DoJSON error")
			}
		})
	})

	t.Run("Execute", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			}))
			defer server.Close()

			client := api.New()

			req := apitest2.NewRequest(http.MethodGet, server.URL)

			resp, err := client.Execute(context.Background(), req)

			if err != nil {
				t.Fatal(err)
			}

			apitest2.AssertOK(t, resp)
		})

		t.Run("BuildHTTPRequest", func(t *testing.T) {
			t.Run("NilBody", func(t *testing.T) {
				server := apitest2.NewServer(t, func(w http.ResponseWriter, r *http.Request) {
					if r.Body != nil {
						// Body should be empty but still valid
						buf := make([]byte, 1)
						n, _ := r.Body.Read(buf)

						if n != 0 {
							t.Fatal("expected empty body")
						}
					}

					w.WriteHeader(http.StatusOK)
				})

				client := api.New()

				req := &transport.Request{
					Method: http.MethodPost,
					Path:   server.URL,
				}

				resp, err := client.Execute(context.Background(), req)
				if err != nil {
					t.Fatal(err)
				}

				apitest2.AssertOK(t, resp)
			})

			t.Run("ReaderError", func(t *testing.T) {
				client := api.New()

				req := &transport.Request{
					Method: http.MethodPost,
					Path:   "https://example.com",
					Body:   failingBody{},
				}

				_, err := client.Execute(context.Background(), req)

				if err == nil {
					t.Fatal("expected body reader error")
				}
			})
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

			resp := apitest2.Execute(t, client, req)

			apitest2.AssertOK(t, resp)
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

			apitest2.AssertOK(t, resp)
		})

		t.Run("Errors", func(t *testing.T) {
			t.Run("Request", func(t *testing.T) {
				client := api.New()

				req := apitest2.NewRequest(http.MethodGet, "://bad-url")

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

				req := apitest2.NewRequest(http.MethodGet, "https://example.com")

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

				req := apitest2.NewRequest(http.MethodGet, "https://example.com")

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

			req := apitest2.NewRequest(http.MethodGet, "https://example.com")

			_, err := client.Execute(context.Background(), req)

			if err == nil {
				t.Fatal("expected credential error")
			}
		})

		t.Run("InvalidVerbMethod", func(t *testing.T) {
			client := api.New()

			req := apitest2.NewRequest("GET\n", "https://example.com")

			_, err := client.Execute(context.Background(), req)
			if err == nil {
				t.Fatal("expected url parse error")
			}
		})
	})
}
