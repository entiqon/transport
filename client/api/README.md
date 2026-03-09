# API Client

Package: `github.com/entiqon/transport/client/api`

The `api` package provides an HTTP transport client for the transport library.

It focuses on executing HTTP requests through a minimal and configurable
interface while remaining independent of credential strategies and
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
4. Apply credentials if configured
5. Execute the request using the configured HTTP client
6. Read and return the transport response

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

| Field | Description |
|------|-------------|
| Connection | Logical connection identifier used by higher-level integrations |
| Method | HTTP method (GET, POST, PUT, DELETE, etc.) |
| Path | Absolute URL or endpoint |
| Headers | Optional HTTP headers |
| Query | Optional query parameters |
| Body | Optional request payload |

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

| Field | Description |
|------|-------------|
| Status | HTTP status code |
| Headers | Response headers |
| Body | Raw response payload |

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

Registers a credential strategy.

```go
api.WithCredential(credential)
```

The credential strategy must implement the `auth.Credential` interface.

Credential implementations are defined in the `auth` and `token`
packages.

---

## Credentials

Credential strategies are independent of the API transport client
and are applied to the request before execution.

Example using an **Access Token header**:

```go
client := api.New(
    api.WithCredential(token.NewAccessToken("X-Access-Token", "token")),
)
```

Example using a **Bearer token**:

```go
client := api.New(
    api.WithCredential(token.NewBearerToken("token")),
)
```

Example using an **API Key**:

```go
client := api.New(
    api.WithCredential(
        token.NewAPIKey("X-API-Key", "token", token.APIKeyHeader),
    ),
)
```

This design keeps the transport layer independent from credential
mechanisms.

---

## License

©️ [Entiqon Labs](https://entiqon.dev)  
[MIT License](../../LICENSE)
