//go:buld test

package apitest

import (
	"context"
	"testing"

	"github.com/entiqon/transport"
)

func Execute(
	t *testing.T,
	client transport.Client,
	req *transport.Request,
) *transport.Response {
	t.Helper()

	resp, err := client.Execute(context.Background(), req)
	if err != nil {
		t.Fatal(err)
	}

	return resp
}
