# auth

Package: `github.com/entiqon/transport/auth`

The `auth` package defines the credential contract used by transport clients.

It provides a minimal interface that allows credential strategies to modify
outgoing HTTP requests before they are executed by a transport client.

The goal of this package is to keep the transport layer independent from
specific authentication mechanisms while allowing different credential
strategies to be plugged in when needed.

---

## Credential Contract

Credential strategies must implement the following interface:

```go
type Credential interface {
    Apply(ctx context.Context, req *http.Request) error
}
```

The `Apply` method is invoked during request execution and allows the
strategy to modify the request, typically by adding authentication data
such as:

- Authorization headers
- API keys
- Access tokens
- Signed request headers

The request will be executed immediately after `Apply` returns.

Implementations should therefore:

- mutate only the provided request
- avoid expensive or blocking operations
- avoid mutating shared state
- be safe for reuse across multiple requests

---

## Example

```go
import (
    "github.com/entiqon/transport/client/api"
    "github.com/entiqon/transport/token"
)

client := api.New(
    api.WithCredential(
        token.NewAccessToken("X-Access-Token", tokenValue),
    ),
)
```

The credential strategy will modify the outgoing HTTP request before it is
executed by the transport client.

---

## Design Goals

The `auth` package focuses exclusively on defining the contract used by
transport clients.

It intentionally avoids implementing authentication mechanisms directly.
Instead, concrete credential strategies are implemented in separate
packages (such as `token`).

This keeps the communication layer small, composable, and independent from
specific authentication strategies.

---

## License

©️ [Entiqon Labs](https://entiqon.dev)  
[MIT License](../../LICENSE)
