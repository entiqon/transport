package errors_test

import (
	"testing"

	"github.com/entiqon/transport/errors"
)

func TestError(t *testing.T) {

	t.Run("CodeOnly", func(t *testing.T) {

		err := &errors.Error{
			Code: "invalid_auth",
		}

		if err.Error() != "invalid_auth" {
			t.Fatalf("unexpected error string: %s", err.Error())
		}

	})

	t.Run("CodeAndMessage", func(t *testing.T) {

		err := &errors.Error{
			Code:    "invalid_auth",
			Message: "invalid authentication configuration",
		}

		expected := "invalid_auth: invalid authentication configuration"

		if err.Error() != expected {
			t.Fatalf("unexpected error string: %s", err.Error())
		}

	})
}
