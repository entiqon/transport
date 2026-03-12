package api_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/entiqon/transport"
	"github.com/entiqon/transport/client/api"
	"github.com/entiqon/transport/credential"
)

// Example_execute demonstrates executing a simple HTTP request using
// the default API client configuration.
func ExampleClient_execute() {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ctx := context.Background()

	client := api.New()

	req := &transport.Request{
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
func ExampleClient_execute_withAccessToken() {

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
		api.WithCredential(
			credential.AccessToken("X-Access-Token", "token"),
		),
	)

	req := &transport.Request{
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

// Example_execute_withBearerToken demonstrates configuring the API client
// with a Bearer token credential and executing a request that requires
// the Authorization header.
func ExampleClient_execute_withBearerToken() {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("Authorization") != "Bearer token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	ctx := context.Background()

	client := api.New(
		api.WithCredential(
			credential.BearerToken("token"),
		),
	)

	req := &transport.Request{
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

// Example_execute_withAPIKey demonstrates configuring the API client
// with an API key credential and executing a request that requires
// the API key to be present in the request headers.
func ExampleClient_Execute_withAPIKey() {

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header.Get("X-API-Key") != "token" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)

	}))
	defer server.Close()

	ctx := context.Background()

	client := api.New(
		api.WithCredential(
			credential.APIKey("X-API-Key", "token", credential.APIKeyHeader),
		),
	)

	req := &transport.Request{
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

func ExampleClient_Execute_withBasePath() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.URL.Path)
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := api.New(
		api.WithBasePath("api"),
	)

	req := &transport.Request{
		Method: http.MethodGet,
		Path:   server.URL + "/users",
	}

	_, _ = client.Execute(context.Background(), req)

	// Output:
	// /api/users
}

func ExampleClient_Execute_withVersion() {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Header.Get("X-API-Version"))
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := api.New(
		api.WithVersion("v1"),
	)

	req := &transport.Request{
		Method: http.MethodGet,
		Path:   server.URL,
	}

	_, _ = client.Execute(context.Background(), req)

	// Output:
	// v1
}
