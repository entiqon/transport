package api

import (
	"context"
	"io"
	"net/http"

	"github.com/entiqon/transport/auth"
)

// api implements a transport client capable of executing HTTP requests.
//
// The client converts a Request into an HTTP request, applies optional
// authentication strategies, performs the request using an http.Client,
// and returns the resulting Response.
type api struct {
	http       *http.Client
	credential auth.Credential
}

// New creates a new API transport client.
//
// The client can be configured through functional options such as
// custom HTTP clients or authentication strategies. If no HTTP client
// is provided, http.DefaultClient is used.
func New(opts ...Option) Client {

	c := &api{
		http: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(c)
	}

	if c.http == nil {
		c.http = http.DefaultClient
	}

	return c
}

// Execute performs the given transport Request.
//
// It validates the request, builds the underlying HTTP request,
// applies authentication if configured, executes the request,
// and returns the resulting Response.
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

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	httpReq, err := http.NewRequestWithContext(ctx, req.Method, req.Path, req.Body)
	if err != nil {
		return nil, err
	}

	for k, v := range req.Headers {
		httpReq.Header.Set(k, v)
	}

	q := httpReq.URL.Query()
	for k, v := range req.Query {
		q.Set(k, v)
	}
	httpReq.URL.RawQuery = q.Encode()

	if c.credential != nil {
		if err := c.credential.Apply(ctx, httpReq); err != nil {
			return nil, err
		}
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	headers := make(map[string]string)
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Response{
		Status:  resp.StatusCode,
		Headers: headers,
		Body:    bodyBytes,
		Raw:     resp,
	}, nil
}
