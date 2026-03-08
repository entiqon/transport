# API Client Documentation

Package: `github.com/entiqon/transport/client/api`

The `api` client provides a minimal HTTP transport implementation for the
transport library. It focuses strictly on executing HTTP requests while
remaining independent from business logic.

---

## Client

Client represents the transport executor.

```go
type Client interface {
    Execute(ctx context.Context, req *Request) (*Response, error)
}
```

The client performs the following steps when executing a request:

1. Validate the request
2. Create an `http.Request`
3. Apply authentication if configured
4. Execute the request using the configured HTTP client
5. Return a minimal transport response

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

| Field | Description |
|------|-------------|
| Method | HTTP method (GET, POST, PUT, DELETE, etc.) |
| Path | Absolute URL or endpoint path |
| Headers | Optional HTTP headers |
| Body | Optional request payload |

---

## Response

Represents the minimal transport response.

```go
type Response struct {
    Status int
}
```

The transport layer intentionally exposes only minimal information.
Applications are responsible for interpreting response payloads.

---

## Client Construction

The API client is created using functional options.

Example:

```go
client := api.New(
    api.WithHTTPClient(http.DefaultClient),
)
```

### Options

#### WithHTTPClient

Provides a custom `http.Client`.

```go
api.WithHTTPClient(httpClient)
```

This allows configuring:

- custom transports
- timeouts
- retries
- observability middleware

#### WithAuth

Registers an authentication strategy.

```go
api.WithAuth(authStrategy)
```

The authentication strategy must implement the `auth.Auth` interface.

---


## License

©️[Entiqon Labs](https://entiqon.dev). [MIT License](../../LICENSE)