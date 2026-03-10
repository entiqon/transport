# auth

Package: `github.com/entiqon/transport/auth`

The `auth` package defines the authentication **contracts** used by the
transport layer.

It separates authentication into three responsibilities:

- **Credential** – modifies outgoing HTTP requests
- **Provider** – resolves credentials dynamically from configuration
- **Refreshable** – optionally allows credentials to be invalidated or refreshed

This separation allows the transport client to remain independent of
specific authentication mechanisms while still supporting dynamic
authentication flows such as OAuth2.

---

## Credential

A **Credential** mutates an outgoing HTTP request in order to apply
authentication data.

```go
type Credential interface {
    Apply(ctx context.Context, req *http.Request) error
}
```

The `Apply` method is invoked during request execution and allows the
credential to modify the request.

Typical use cases include:

- Authorization headers
- API keys
- Bearer tokens
- HMAC signatures
- JWT headers

Credential implementations are provided by the
`github.com/entiqon/transport/credential` package.

---

## Provider

A **Provider** resolves a credential from configuration.

```go
type Provider interface {
    Resolve(ctx context.Context, cfg config.Auth) (Credential, error)
}
```

Providers enable dynamic authentication flows where credentials must be
retrieved or refreshed before being applied to a request.

Examples include:

- OAuth2 token resolution
- OAuth2 refresh flows
- external credential services

Provider implementations are provided by the
`github.com/entiqon/transport/provider` package.

---

## Refreshable

Some authentication providers maintain internal credential state such as
cached access tokens.

Providers that support forced credential renewal may implement the
`Refreshable` interface.

```go
type Refreshable interface {
    Refresh(ctx context.Context) error
}
```

The `Refresh` method invalidates any cached credentials so that the next
call to `Provider.Resolve` retrieves fresh credentials.

This is typically used when a transport client receives an
`HTTP 401 Unauthorized` response and needs to force credential renewal.

Not all providers implement this interface.

---

## Static Credential Example

```go
client := api.New(
    api.WithCredential(
        credential.BearerToken("token"),
    ),
)
```

The credential strategy modifies the outgoing HTTP request before it is
executed by the transport client.

---

## Provider-based Authentication Example

```go
authConfig := config.Auth{
    Kind: "oauth2",
    OAuth2: &config.OAuth2{
        GrantType:    "client_credentials",
        TokenURL:     "https://auth.example.com/oauth2/token",
        ClientID:     "client",
        ClientSecret: "secret",
    },
}

client := api.New(
    api.WithAuthProvider(
        provider.OAuth2(http.DefaultClient),
        authConfig,
    ),
)
```

The provider resolves the credential dynamically before the request is
executed.

---

## Design Goals

The `auth` package intentionally defines **contracts only**.

Concrete implementations are located in separate packages:

| Package | Responsibility |
|--------|----------------|
| `credential` | request authentication strategies |
| `provider` | credential resolution mechanisms |

This design keeps the transport layer small, composable, and extensible.

---

## License

© Entiqon Labs  
Released under the MIT License.
