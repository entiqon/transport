package token

import (
	"context"
	"net/http"
)

type Bearer struct {
	Provider Provider
}

func (b Bearer) Apply(ctx context.Context, req *http.Request) error {
	t, err := b.Provider.Token(ctx)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", "Bearer "+t)
	return nil
}
