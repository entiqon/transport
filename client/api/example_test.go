package api_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/entiqon/transport/client/api"
	"github.com/entiqon/transport/token"
)

// Example_execute demonstrates executing a simple HTTP request using
// the default API client configuration.
func Example_execute() {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx := context.Background()

	client := api.New()

	req := &api.Request{
		Method: "GET",
		Path:   server.URL,
	}

	resp, err := client.Execute(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)

	// Output:
	// 200
}

// Example_execute_withAccessToken demonstrates configuring the API client
// with an access token authentication strategy and executing a request
// that requires the token to be present in the request headers.
func Example_execute_withAccessToken() {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("X-Access-Token") != "token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	ctx := context.Background()

	client := api.New(
		api.WithAuth(
			token.NewAccessToken("X-Access-Token", "token"),
		),
	)

	req := &api.Request{
		Method: "GET",
		Path:   server.URL,
	}

	resp, err := client.Execute(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)

	// Output:
	// 200
}
