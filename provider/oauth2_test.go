package provider_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	"github.com/entiqon/transport/config"
	"github.com/entiqon/transport/provider"
)

type oauthServer struct {
	srv   *httptest.Server
	calls int32
}

func newOAuthServer(
	handler func(call int32, r *http.Request) map[string]any,
) *oauthServer {

	var calls int32

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		n := atomic.AddInt32(&calls, 1)

		resp := handler(n, r)

		_ = json.NewEncoder(w).Encode(resp)

	}))

	return &oauthServer{
		srv: srv,
	}
}

func (s *oauthServer) URL() string {
	return s.srv.URL
}

func (s *oauthServer) Client() *http.Client {
	return s.srv.Client()
}

func (s *oauthServer) Close() {
	s.srv.Close()
}

type failingTransport struct{}

func (f failingTransport) RoundTrip(*http.Request) (*http.Response, error) {
	return nil, errors.New("network error")
}

func TestOAuth2Provider(t *testing.T) {

	t.Run("Request", func(t *testing.T) {

		t.Run("InvalidURL", func(t *testing.T) {

			p := provider.OAuth2(http.DefaultClient)

			cfg := config.Auth{
				Kind: "oauth2",
				OAuth2: &config.OAuth2{
					GrantType:    "client_credentials",
					TokenURL:     "://invalid-url", // forces NewRequestWithContext error
					ClientID:     "client",
					ClientSecret: "secret",
				},
			}

			_, err := p.Resolve(context.Background(), cfg)
			if err == nil {
				t.Fatal("expected request creation error")
			}
		})

		t.Run("InvalidJSON", func(t *testing.T) {

			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("invalid-json"))
			}))
			defer srv.Close()

			p := provider.OAuth2(srv.Client())

			cfg := config.Auth{
				Kind: "oauth2",
				OAuth2: &config.OAuth2{
					GrantType:    "client_credentials",
					TokenURL:     srv.URL,
					ClientID:     "client",
					ClientSecret: "secret",
				},
			}

			_, err := p.Resolve(context.Background(), cfg)
			if err == nil {
				t.Fatal("expected json decode error")
			}
		})

		t.Run("HTTPError", func(t *testing.T) {

			client := &http.Client{
				Transport: failingTransport{},
			}

			p := provider.OAuth2(client)

			cfg := config.Auth{
				Kind: "oauth2",
				OAuth2: &config.OAuth2{
					GrantType:    "client_credentials",
					TokenURL:     "https://example.com/token",
					ClientID:     "client",
					ClientSecret: "secret",
				},
			}

			_, err := p.Resolve(context.Background(), cfg)
			if err == nil {
				t.Fatal("expected http error")
			}
		})
	})

	t.Run("Resolve", func(t *testing.T) {

		t.Run("RefreshToken", func(t *testing.T) {
			srv := newOAuthServer(func(call int32, r *http.Request) map[string]any {
				return map[string]any{
					"access_token": "token-1",
					"expires_in":   3600,
				}
			})
			defer srv.Close()

			oauth2 := provider.OAuth2(srv.Client())

			cfg := config.Auth{
				Kind: "oauth2",
				OAuth2: &config.OAuth2{
					GrantType:    "refresh_token",
					TokenURL:     srv.URL(),
					ClientID:     "client",
					ClientSecret: "secret",
					RefreshToken: "refresh",
				},
			}

			cred, err := oauth2.Resolve(context.Background(), cfg)
			if err != nil {
				t.Fatal(err)
			}

			if cred == nil {
				t.Fatal("expected credential")
			}
		})

		t.Run("ClientCredentials", func(t *testing.T) {

			srv := newOAuthServer(func(call int32, r *http.Request) map[string]any {

				if err := r.ParseForm(); err != nil {
					t.Fatal(err)
				}

				if r.Form.Get("grant_type") != "client_credentials" {
					t.Fatal("expected client_credentials grant")
				}

				return map[string]any{
					"access_token": "token-cc",
					"expires_in":   3600,
				}
			})
			defer srv.Close()

			oauth2 := provider.OAuth2(srv.Client())

			cfg := config.Auth{
				Kind: "oauth2",
				OAuth2: &config.OAuth2{
					GrantType:    "client_credentials",
					TokenURL:     srv.URL(),
					ClientID:     "client",
					ClientSecret: "secret",
				},
			}

			cred, err := oauth2.Resolve(context.Background(), cfg)
			if err != nil {
				t.Fatal(err)
			}

			if cred == nil {
				t.Fatal("expected credential")
			}
		})
	})

	t.Run("Cache", func(t *testing.T) {

		var calls int32

		srv := newOAuthServer(func(call int32, r *http.Request) map[string]any {

			atomic.AddInt32(&calls, 1)

			return map[string]any{
				"access_token": "token-1",
				"expires_in":   3600,
			}
		})
		defer srv.Close()

		oauth2 := provider.OAuth2(srv.Client())

		cfg := config.Auth{
			Kind: "oauth2",
			OAuth2: &config.OAuth2{
				GrantType:    "refresh_token",
				TokenURL:     srv.URL(),
				ClientID:     "client",
				ClientSecret: "secret",
				RefreshToken: "refresh",
			},
		}

		ctx := context.Background()

		_, err := oauth2.Resolve(ctx, cfg)
		if err != nil {
			t.Fatal(err)
		}

		_, err = oauth2.Resolve(ctx, cfg)
		if err != nil {
			t.Fatal(err)
		}

		if atomic.LoadInt32(&calls) != 1 {
			t.Fatalf("expected 1 token request, got %d", calls)
		}
	})

	t.Run("Refresh", func(t *testing.T) {

		t.Run("Success", func(t *testing.T) {
			t.Run("Refresh", func(t *testing.T) {

				srv := newOAuthServer(func(call int32, r *http.Request) map[string]any {
					return map[string]any{
						"access_token": "token-1",
						"expires_in":   3600,
					}
				})
				defer srv.Close()

				oauth2 := provider.OAuth2(srv.Client())

				cfg := config.Auth{
					Kind: "oauth2",
					OAuth2: &config.OAuth2{
						GrantType:    "refresh_token",
						TokenURL:     srv.URL(),
						ClientID:     "client",
						ClientSecret: "secret",
						RefreshToken: "refresh",
					},
				}

				ctx := context.Background()

				_, err := oauth2.Resolve(ctx, cfg)
				if err != nil {
					t.Fatal(err)
				}

				if r, ok := oauth2.(interface{ Refresh(context.Context) error }); ok {
					if err := r.Refresh(ctx); err != nil {
						t.Fatal(err)
					}
				}

				_, err = oauth2.Resolve(ctx, cfg)
				if err != nil {
					t.Fatal(err)
				}
			})
		})

		t.Run("OnExpired", func(t *testing.T) {

			var calls int32

			srv := newOAuthServer(func(call int32, r *http.Request) map[string]any {

				n := atomic.AddInt32(&calls, 1)

				token := "token-1"
				if n > 1 {
					token = "token-2"
				}

				return map[string]any{
					"access_token": token,
					"expires_in":   1,
				}
			})
			defer srv.Close()

			oauth2 := provider.OAuth2(srv.Client())

			cfg := config.Auth{
				Kind: "oauth2",
				OAuth2: &config.OAuth2{
					GrantType:    "refresh_token",
					TokenURL:     srv.URL(),
					ClientID:     "client",
					ClientSecret: "secret",
					RefreshToken: "refresh",
				},
			}

			ctx := context.Background()

			_, err := oauth2.Resolve(ctx, cfg)
			if err != nil {
				t.Fatal(err)
			}

			time.Sleep(1100 * time.Millisecond)

			_, err = oauth2.Resolve(ctx, cfg)
			if err != nil {
				t.Fatal(err)
			}

			if atomic.LoadInt32(&calls) != 2 {
				t.Fatalf("expected refresh call")
			}
		})

		t.Run("Rotation", func(t *testing.T) {

			srv := newOAuthServer(func(call int32, r *http.Request) map[string]any {
				return map[string]any{
					"access_token":  "token-1",
					"refresh_token": "refresh-2",
					"expires_in":    3600,
				}
			})
			defer srv.Close()

			oauth2 := provider.OAuth2(srv.Client())

			cfg := config.Auth{
				Kind: "oauth2",
				OAuth2: &config.OAuth2{
					GrantType:    "refresh_token",
					TokenURL:     srv.URL(),
					ClientID:     "client",
					ClientSecret: "secret",
					RefreshToken: "refresh-1",
				},
			}

			_, err := oauth2.Resolve(context.Background(), cfg)
			if err != nil {
				t.Fatal(err)
			}

			if cfg.OAuth2.RefreshToken != "refresh-2" {
				t.Fatal("refresh token was not rotated")
			}
		})
	})

	t.Run("Config", func(t *testing.T) {

		t.Run("ClientCredentials", func(t *testing.T) {

			t.Run("WithScope", func(t *testing.T) {

				srv := newOAuthServer(func(call int32, r *http.Request) map[string]any {

					if err := r.ParseForm(); err != nil {
						t.Fatal(err)
					}

					if r.Form.Get("scope") != "read:orders" {
						t.Fatalf("expected scope to be sent")
					}

					return map[string]any{
						"access_token": "token-cc",
						"expires_in":   3600,
					}
				})
				defer srv.Close()

				oauth2 := provider.OAuth2(srv.Client())

				cfg := config.Auth{
					Kind: "oauth2",
					OAuth2: &config.OAuth2{
						GrantType:    "client_credentials",
						TokenURL:     srv.URL(),
						ClientID:     "client",
						ClientSecret: "secret",
						Scope:        "read:orders",
					},
				}

				cred, err := oauth2.Resolve(context.Background(), cfg)
				if err != nil {
					t.Fatal(err)
				}

				if cred == nil {
					t.Fatal("expected credential")
				}
			})
		})

		t.Run("Invalid", func(t *testing.T) {

			t.Run("Self", func(t *testing.T) {
				oauth2 := provider.OAuth2(nil)

				cfg := config.Auth{
					Kind: "oauth2",
				}

				_, err := oauth2.Resolve(context.Background(), cfg)

				if err == nil {
					t.Fatal("expected error")
				}
			})

			t.Run("GrantType", func(t *testing.T) {

				oauth2 := provider.OAuth2(nil)

				cfg := config.Auth{
					Kind: "oauth2",
					OAuth2: &config.OAuth2{
						GrantType: "invalid",
					},
				}

				_, err := oauth2.Resolve(context.Background(), cfg)

				if err == nil {
					t.Fatal("expected error")
				}
			})
		})
	})
}
