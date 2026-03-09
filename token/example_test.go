package token_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entiqon/transport/token"
)

func ExampleNewAccessToken() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := token.NewAccessToken("X-Access-Token", "abc123")

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("X-Access-Token"))

	// Output:
	// abc123
}

func ExampleNewBearerToken() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := token.NewBearerToken("abc123")

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("Authorization"))

	// Output:
	// Bearer abc123
}

func ExampleNewAPIKey_header() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := token.NewAPIKey("X-API-Key", "abc123", token.APIKeyHeader)

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("X-API-Key"))

	// Output:
	// abc123
}

func ExampleNewAPIKey_query() {

	req, _ := http.NewRequest("GET", "https://example.com/resource", nil)

	cred := token.NewAPIKey("api_key", "abc123", token.APIKeyQuery)

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.URL.Query().Get("api_key"))

	// Output:
	// abc123
}
