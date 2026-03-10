package provider_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/entiqon/transport"
	"github.com/entiqon/transport/client/api"
	"github.com/entiqon/transport/config"
	"github.com/entiqon/transport/provider"
)

func ExampleOAuth2() {

	// Mock OAuth2 token endpoint
	tokenServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		_, _ = fmt.Fprint(w, `{
					"access_token": "example-token",
					"expires_in": 3600,
					"token_type": "Bearer"
				}`)
	}))
	defer tokenServer.Close()

	// Mock API endpoint
	apiServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")

		if auth != "Bearer example-token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	}))
	defer apiServer.Close()

	authConfig := config.Auth{
		Kind: "oauth2",
		OAuth2: &config.OAuth2{
			GrantType:    "refresh_token",
			TokenURL:     tokenServer.URL,
			ClientID:     "client-id",
			ClientSecret: "client-secret",
			RefreshToken: "refresh-token",
		},
	}

	ouath2 := provider.OAuth2(nil)

	client := api.New(
		api.WithAuthProvider(ouath2, authConfig),
	)

	req := &transport.Request{
		Method: "GET",
		Path:   apiServer.URL,
	}

	resp, err := client.Execute(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(resp.Status)

	// Output:
	// 200
}
