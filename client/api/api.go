package api

import (
	"context"
	"net/http"

	"github.com/entiqon/transport/auth"
)

type api struct {
	http *http.Client
	auth auth.Auth
}

// New creates a new API client.
func New(opts ...Option) Client {

	c := &api{
		http: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	// validation
	if c.http == nil {
		c.http = http.DefaultClient
	}

	return c
}

func (c *api) Execute(ctx context.Context, req *Request) (*Response, error) {

	if req == nil {
		return nil, InvalidRequestError
	}

	if req.Method == "" {
		return nil, MissingMethodError
	}

	if req.Path == "" {
		return nil, MissingPathError
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.Path, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	if c.auth != nil {
		if err := c.auth.Apply(ctx, httpReq); err != nil {
			return nil, err
		}
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return &Response{
		Status: resp.StatusCode,
	}, nil
}
