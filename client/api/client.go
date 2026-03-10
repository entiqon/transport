package api

import (
	"context"
	"io"
	"net/http"

	"github.com/entiqon/transport"
	"github.com/entiqon/transport/auth"
	"github.com/entiqon/transport/config"
)

// api implements a transport client capable of executing HTTP requests.
//
// The client converts a Request into an HTTP request, applies optional
// authentication strategies, performs the request using an http.Client,
// and returns the resulting Response.
type client struct {
	http       *http.Client
	credential auth.Credential
	provider   auth.Provider
	config     config.Auth
}

// New creates a new API transport client.
//
// The client can be configured through functional options such as
// custom HTTP clients or authentication strategies. If no HTTP client
// is provided, http.DefaultClient is used.
func New(opts ...Option) transport.Client {

	c := &client{
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
func (c *client) Execute(ctx context.Context, req *transport.Request) (*transport.Response, error) {

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

	httpReq, err := c.buildHTTPRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	if err := c.applyCredential(ctx, httpReq); err != nil {
		return nil, err
	}

	resp, err := c.http.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return c.buildResponse(resp)
}

// applyCredential resolves and applies authentication credentials
// to the provided HTTP request if authentication is configured.
func (c *client) applyCredential(
	ctx context.Context,
	req *http.Request,
) error {

	cred, err := c.resolveCredential(ctx)
	if err != nil {
		return err
	}

	if cred == nil {
		return nil
	}

	return cred.Apply(ctx, req)
}

// buildHTTPRequest constructs an HTTP request from the transport Request.
func (c *client) buildHTTPRequest(
	ctx context.Context,
	req *transport.Request,
) (*http.Request, error) {

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

	return httpReq, nil
}

// buildResponse converts an HTTP response into a transport Response.
func (c *client) buildResponse(resp *http.Response) (*transport.Response, error) {

	headers := make(map[string]string, len(resp.Header))

	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &transport.Response{
		Status:  resp.StatusCode,
		Headers: headers,
		Body:    bodyBytes,
		Raw:     resp,
	}, nil
}

func (c *client) resolveCredential(
	ctx context.Context,
) (auth.Credential, error) {

	if c.credential != nil {
		return c.credential, nil
	}

	if c.provider == nil {
		return nil, nil
	}

	if c.config.Kind == "" {
		return nil, config.InvalidAuthConfigError
	}

	return c.provider.Resolve(ctx, c.config)
}
