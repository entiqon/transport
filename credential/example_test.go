package credential_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entiqon/transport/credential"
)

func ExampleAccessToken() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := credential.AccessToken("X-Access-Token", "abc123")

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("X-Access-Token"))

	// Output:
	// abc123
}

func ExampleBearerToken() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := credential.BearerToken("abc123")

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("Authorization"))

	// Output:
	// Bearer abc123
}

func ExampleAPIKey_header() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := credential.APIKey("X-API-Key", "abc123", credential.APIKeyHeader)

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("X-API-Key"))

	// Output:
	// abc123
}

func ExampleAPIKey_query() {

	req, _ := http.NewRequest("GET", "https://example.com/resource", nil)

	cred := credential.APIKey("api_key", "abc123", credential.APIKeyQuery)

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.URL.Query().Get("api_key"))

	// Output:
	// abc123
}

func ExampleBasic() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := credential.Basic("user", "pass")

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("Authorization")[:5])

	// Output:
	// Basic
}

func ExampleJWT() {

	req, _ := http.NewRequest("GET", "https://example.com", nil)

	cred := credential.JWT("Authorization", "jwt_token")

	_ = cred.Apply(context.Background(), req)

	fmt.Println(req.Header.Get("Authorization"))

	// Output:
	// Bearer jwt_token
}
