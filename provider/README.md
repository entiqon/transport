# providers

Package: `github.com/entiqon/transport/providers`

The `providers` package contains implementations of the `auth.Provider`
interface.

Providers resolve authentication configuration into credentials that can be
applied to outgoing HTTP requests by transport clients.

Each provider is responsible for handling the lifecycle of its authentication
mechanism, including retrieving credentials, refreshing tokens, and managing
expiration when necessary.

---

## OAuth2Provider

`OAuth2Provider` resolves OAuth2 authentication configuration into a
`BearerToken` credential.

The provider performs the following responsibilities:

* retrieves OAuth2 access tokens from the configured token endpoint
* refreshes tokens when they expire
* caches tokens between requests
* supports refresh token rotation when provided by the authorization server

Supported OAuth2 grant types currently include:

* `refresh_token`
* `client_credentials`

---

## Example

```go
import (
    "context"

    "github.com/entiqon/transport/client/api"
    "github.com/entiqon/transport/config"
    "github.com/entiqon/transport/providers"
)

authConfig := config.Auth{
    Kind: "oauth2",
    OAuth2: &config.OAuth2{
        GrantType:    "refresh_token",
        TokenURL:     "https://auth.example.com/token",
        ClientID:     "...",
        ClientSecret: "...",
        RefreshToken: "...",
    },
}

provider := providers.NewOAuth2Provider(nil)

client := api.New(
    api.WithAuthProvider(provider, authConfig),
)

resp, err := client.Execute(context.Background(), req)
```

During request execution the provider resolves a credential which is then
applied to the outgoing HTTP request by the transport client.

---

## Design Goals

Providers implement authentication flows while keeping the transport client
independent from specific authentication mechanisms.

This separation allows authentication strategies to evolve independently from
the HTTP transport layer.

---

## License

©️ [Entiqon Labs](https://entiqon.dev)
[MIT License](../../LICENSE)
