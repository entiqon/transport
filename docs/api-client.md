# API Client

The `client/api` package provides a minimal HTTP transport client.

It focuses strictly on executing HTTP requests while remaining
independent from application logic.

## Client

Client represents the HTTP transport executor.

```go
type Client interface {
    Execute(ctx context.Context, req *Request) (*Response, error)
}
```

Execution flow:

1. Validate the request
2. Build an `http.Request`
3. Apply authentication if configured
4. Execute through the configured HTTP client
5. Return the transport response

---

## Request

Represents a transport request.

```go
type Request struct {
    Method  string
    Path    string
    Headers map[string]string
    Body    any
}
```

Fields:

- **Method** – HTTP method (GET, POST, etc.)
- **Path** – Absolute URL or endpoint path
- **Headers** – Optional HTTP headers
- **Body** – Optional payload

---

## Response

Minimal response returned by the transport.

```go
type Response struct {
    Status int
}
```

The response intentionally exposes only transport-level data.

Applications are responsible for interpreting response payloads.

---

## Options

The API client uses the functional options pattern.

Example:

```go
client := api.New(
    api.WithHTTPClient(http.DefaultClient),
    api.WithAuth(auth),
)
```

### WithHTTPClient

Allows providing a custom `http.Client`.

Useful for:

- retry logic
- custom transports
- timeouts
- observability middleware

### WithAuth

Registers an authentication strategy implementing the `auth.Auth`
interface.
