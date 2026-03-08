package token

import "context"

type token struct {
	Value string
}

type Provider interface {
	Token(ctx context.Context) (string, error)
}

func (t token) Token(ctx context.Context) (string, error) {
	return t.Value, nil
}
