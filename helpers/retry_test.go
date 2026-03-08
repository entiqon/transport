package helpers_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/entiqon/transport/helpers"
)

func TestRetry(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		t.Parallel()

		t.Run("FirstAttempt", func(t *testing.T) {
			err := helpers.Retry(context.Background(), 3, time.Nanosecond, func() error {
				return nil
			})

			if err != nil {
				t.Fatalf("expected nil error, got %v", err)
			}
		})

		t.Run("SecondAttempt", func(t *testing.T) {
			attempts := 0

			err := helpers.Retry(context.Background(), 3, time.Nanosecond, func() error {
				attempts++
				if attempts < 2 {
					return errors.New("fail")
				}
				return nil
			})

			if err != nil {
				t.Fatal(err)
			}

			if attempts != 2 {
				t.Fatalf("expected 2 attempts, got %d", attempts)
			}
		})
	})

	t.Run("ExhaustAttempts", func(t *testing.T) {
		attempts := 0

		err := helpers.Retry(context.Background(), 3, time.Nanosecond, func() error {
			attempts++
			return errors.New("fail")
		})

		if err == nil {
			t.Fatal("expected error")
		}

		if attempts != 3 {
			t.Fatalf("expected 3 attempts, got %d", attempts)
		}
	})

	t.Run("ContextCanceled", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		err := helpers.Retry(ctx, 3, time.Nanosecond, func() error {
			return errors.New("fail")
		})

		if err == nil {
			t.Fatal("expected context error")
		}
	})
}
