package helpers

import (
	"context"
	"time"
)

// Retry executes a function multiple times with delay.
func Retry(ctx context.Context, attempts int, delay time.Duration, fn func() error) error {
	var err error

	for i := 0; i < attempts; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(delay):
		}
	}

	return err
}
