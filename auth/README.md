# auth

Package: `github.com/entiqon/transport/auth`

`auth` defines the authentication contract used by transport clients.

This package provides a minimal interface that allows authentication
strategies to modify outgoing HTTP requests before they are executed by a
transport client.

The goal is to keep the transport layer independent from specific
authentication mechanisms while allowing different strategies to be plugged
in when needed.

---

## Auth Contract

Authentication strategies must implement the following interface:

```go
type Auth interface {
    Apply(ctx context.Context, req *http.Request) error
}
```

The `Apply` method is invoked during request execution and allows the
strategy to modify the request, typically by adding headers such as:

- Authorization
- API keys
- access tokens

---

## Design Goals

The `auth` package focuses only on defining the contract used by transport
clients. It intentionally avoids implementing authentication mechanisms.

This keeps the communication layer small, composable, and independent from
specific authentication strategies.

Authentication implementations are defined in separate packages.

---

## License

©️ [Entiqon Labs](https://entiqon.dev)  
[MIT License](../../LICENSE)