package api_test

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entiqon/transport/client/api"
)

// Example_apiClient demonstrates creating and using the API client.
func Example_apiClient() {

	ctx := context.Background()

	client := api.New(
		api.WithHTTPClient(http.DefaultClient),
	)

	req := &api.Request{
		Method: "GET",
		Path:   "https://example.com",
		Headers: map[string]string{
			"Accept": "application/json",
		},
	}

	resp, err := client.Execute(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)
}
