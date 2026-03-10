# entiqon/transport

[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/transport.svg)](https://pkg.go.dev/github.com/entiqon/transport)
[![CI](https://github.com/entiqon/transport/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/transport/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/entiqon/transport/branch/main/graph/badge.svg)](https://codecov.io/gh/entiqon/transport)
[![Go Report Card](https://goreportcard.com/badge/github.com/entiqon/transport)](https://goreportcard.com/report/github.com/entiqon/transport)
[![Latest Release](https://img.shields.io/github/v/release/entiqon/transport)](https://github.com/entiqon/transport/releases)
[![License](https://img.shields.io/github/license/entiqon/transport)](https://github.com/entiqon/transport/blob/main/LICENSE)

`transport` is a minimal Go library providing reusable primitives
for executing requests across different communication transports.

The library focuses strictly on the **communication layer**, allowing
applications to interact with external systems through a unified
transport abstraction.

It is designed to remain **small, composable, and transport-focused**,
leaving orchestration, retries, and domain logic to the consuming
application.

---

## Architecture

The library is organized into small composable packages:

```
transport (core primitives)
├── Client
├── Request
└── Response
      ↑
client/api (HTTP implementation)
      ↑
auth (authentication contracts)
      ↑
credential (request mutation strategies)
      ↑
provider (credential resolution)
```

This layering ensures:

- transport execution is independent from authentication
- credentials remain pluggable
- providers can dynamically resolve credentials

---

## Quick Example

```go
package main

import (
    "context"
    "fmt"
    "net/http"

    "github.com/entiqon/transport"
    "github.com/entiqon/transport/client/api"
)

func main() {

    ctx := context.Background()

    client := api.New(
        api.WithHTTPClient(http.DefaultClient),
    )

    req := &transport.Request{
        Method: "GET",
        Path:   "https://example.com",
    }

    resp, err := client.Execute(ctx, req)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp.Status)
}
```

---

## Credential Strategies

Credential strategies mutate outgoing requests before execution.

### Bearer Token

```go
client := api.New(
    api.WithCredential(
        credential.BearerToken("token"),
    ),
)
```

### API Key

```go
client := api.New(
    api.WithCredential(
        credential.APIKey("X-API-Key", "key", credential.APIKeyHeader),
    ),
)
```

### Access Token Header

```go
client := api.New(
    api.WithCredential(
        credential.AccessToken("X-Access-Token", "token"),
    ),
)
```

### Basic Authentication

```go
client := api.New(
    api.WithCredential(
        credential.Basic("user", "password"),
    ),
)
```

### JWT

```go
client := api.New(
    api.WithCredential(
        credential.JWT("Authorization", jwtToken),
    ),
)
```

### HMAC Signing

```go
client := api.New(
    api.WithCredential(
        credential.HMAC("api-key", "secret"),
    ),
)
```

---

## Credential Providers

Credential providers resolve credentials dynamically from configuration.

Example using an OAuth2 provider:

```go
client := api.New(
    api.WithAuthProvider(
        provider.OAuth2(http.DefaultClient),
        authConfig,
    ),
)
```

Providers may implement automatic credential refresh when tokens expire.

---

## Design Goals

The transport library focuses exclusively on **communication concerns**.

It intentionally avoids:

- business logic
- retry orchestration
- domain transformations
- workflow coordination

These responsibilities belong to the consuming application.

---

## License

Originally created by **Isidro A. Lopez G.**  
Maintained by **Entiqon Labs**

MIT License
