# API Client

Package: `github.com/entiqon/transport/client/api`

The `api` package provides an HTTP transport client for the transport
library.

It focuses on executing HTTP requests through a minimal and configurable
interface while remaining independent of authentication strategies and
application-specific behavior.

---

## Client

`Client` represents a transport executor.

```go
type Client interface {
    Execute(ctx context.Context, req *Request) (*Response, error)
}
```

When executing a request, the client performs the following steps:

1. Validate the request
2. Construct an `http.Request`
3. Apply authentication if configured
4. Execute the request using the configured HTTP client
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

| Field   | Description                                |
|---------|--------------------------------------------|
| Method  | HTTP method (GET, POST, PUT, DELETE, etc.) |
| Path    | Absolute URL or endpoint                   |
| Headers | Optional HTTP headers                      |
| Body    | Optional request payload                   |

---

## Response

Represents the transport response.

```go
type Response struct {
    Status int
}
```

The API client intentionally exposes a minimal response structure.
Additional response processing can be implemented by the caller.

---

## Client Construction

The API client is created using functional options.

```go
client := api.New(
    api.WithHTTPClient(http.DefaultClient),
)
```

---

## Options

### WithHTTPClient

Provides a custom `http.Client`.

```go
api.WithHTTPClient(httpClient)
```

This allows configuration of:

- custom transports
- timeouts
- connection behavior
- instrumentation

---

### WithAuth

Registers an authentication strategy.

```go
api.WithAuth(authStrategy)
```

The authentication strategy must implement the `auth.Auth` interface.

Authentication implementations are defined in the `auth` package.

---

## Authentication

Authentication strategies are independent of the API transport client.
They are applied to the HTTP request before execution.

Example:

```go
client := api.New(
    api.WithAuth(auth.NewAccessToken("X-Access-Token", "token")),
)
```

---

## License

©️ [Entiqon Labs](https://entiqon.dev)  
[MIT License](../../LICENSE)
