//go:buld test

package apitest

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()

	server := httptest.NewServer(handler)

	t.Cleanup(func() {
		server.Close()
	})

	return server
}
