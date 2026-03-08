package auth

import (
	"context"
	"net/http"

	"github.com/entiqon/transport/auth/token"
)

type Bearer struct {
	Provider token.Provider
}

func (b Bearer) Apply(ctx context.Context, req *http.Request) error {
	t, err := b.Provider.Token(ctx)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+t)
	return nil
}
