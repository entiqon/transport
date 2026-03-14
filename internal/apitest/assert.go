//go:buld test

package apitest

import (
	"testing"

	"github.com/entiqon/transport"
)

func AssertOK(t *testing.T, resp *transport.Response) {
	t.Helper()

	if !resp.OK() {
		t.Fatalf("unexpected status: %d", resp.Status)
	}
}
