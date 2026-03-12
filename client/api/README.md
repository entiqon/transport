
# API Client

Package: `github.com/entiqon/transport/client/api`

The `api` package provides an HTTP transport client for the transport library.

It focuses on executing HTTP requests through a minimal and configurable
interface while remaining independent of authentication mechanisms and
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
3. Apply request headers and query parameters
4. Resolve credentials from an auth provider (if configured)
5. Apply credentials to the request
6. Execute the request using the configured HTTP client
7. Read and return the transport response

---

## Request

Represents a transport request.

```go
type Request struct {
    Connection string
    Method     string
    Path       string
    Headers    map[string]string
    Query      map[string]string
    Body       io.Reader
}
```

| Field      | Description                                                     |
|------------|-----------------------------------------------------------------|
| Connection | Logical connection identifier used by higher-level integrations |
| Method     | HTTP method (GET, POST, PUT, DELETE, etc.)                      |
| Path       | Absolute URL or endpoint                                        |
| Headers    | Optional HTTP headers                                           |
| Query      | Optional query parameters                                       |
| Body       | Optional request payload                                        |

The `Body` field accepts any `io.Reader`, allowing flexible payload
streaming such as JSON encoders, byte buffers, or files.

---

## Response

Represents the transport response.

```go
type Response struct {
    Status  int
    Headers map[string]string
    Body    []byte
}
```

| Field   | Description          |
|---------|----------------------|
| Status  | HTTP status code     |
| Headers | Response headers     |
| Body    | Raw response payload |

The API client intentionally returns the raw payload so that callers
can perform custom decoding or processing.

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
- connection pooling
- instrumentation
- retry logic

---

### WithCredential

Registers a static credential strategy.

```go
api.WithCredential(credential)
```

The credential must implement the `auth.Credential` interface and
is applied to the request before execution.

---

### WithAuthProvider

Registers an authentication provider capable of dynamically resolving
credentials.

```go
api.WithAuthProvider(provider, config)
```

Providers implement the `auth.Provider` interface and resolve
credentials before the request is executed.

Some providers may support credential renewal by implementing
`auth.Refreshable`.

---

## Authentication

The API client does not implement authentication directly.

Authentication is applied through either:

- **Credential strategies**, which statically apply credentials to a request
- **Authentication providers**, which dynamically resolve credentials
  before execution

This design keeps the transport layer independent from authentication
mechanisms and allows integrations to plug in their own credential
implementations.

---

## License

©️ [Entiqon Labs](https://entiqon.dev)  
[MIT License](../../LICENSE)
