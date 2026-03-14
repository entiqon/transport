package helpers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewTestServer(
	t *testing.T,
	handler func(http.ResponseWriter, *http.Request),
) *httptest.Server {

	t.Helper()

	s := httptest.NewServer(http.HandlerFunc(handler))

	t.Cleanup(func() {
		s.Close()
	})

	return s
}
