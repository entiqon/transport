# entiqon/transport

[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/transport.svg)](https://pkg.go.dev/github.com/entiqon/transport)
[![CI](https://github.com/entiqon/transport/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/transport/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/entiqon/transport/branch/main/graph/badge.svg)](https://codecov.io/gh/entiqon/transport)
[![Go Report Card](https://goreportcard.com/badge/github.com/entiqon/transport)](https://goreportcard.com/report/github.com/entiqon/transport)
[![Latest Release](https://img.shields.io/github/v/release/entiqon/transport)](https://github.com/entiqon/transport/releases)
[![License](https://img.shields.io/github/license/entiqon/transport)](https://github.com/entiqon/transport/blob/main/LICENSE)
![Documentation](https://img.shields.io/badge/docs-coming--soon-lightgrey)

`transport` is a minimal Go library that provides reusable primitives
for executing requests across different communication transports.

The library focuses strictly on the **communication layer**, allowing
applications to interact with external systems through a unified
transport abstraction. It is designed to remain **small, composable,
and transport-focused**, leaving business logic, orchestration, and
data transformation to the application layer.

---

## Why transport?

Applications often interact with external systems through different
communication channels such as APIs, SFTP servers, or webhooks.

`transport` provides a small set of primitives that make those interactions consistent while keeping application logic
independent of the underlying communication mechanism.

---

## Goals

- Provide a minimal transport client abstraction
- Support multiple communication channels (API, SFTP, Webhooks, etc.)
- Offer reusable authentication primitives
- Keep configuration models simple and transport-specific
- Remain lightweight and easy to integrate

---

## Example

``` go
package main

import (
    "context"
    "fmt"

    "github.com/entiqon/transport/client"
)

func main() {
    ctx := context.Background()

    req := &client.Request{
        Connection: "example-api",
        Method:     "POST",
        Path:       "/orders",
        Headers: map[string]string{
            "Content-Type": "application/json",
        },
        Body: map[string]any{
            "id":    123,
            "total": 99.99,
        },
    }

    // client implementation provided by the application
    var c client.Client
    

    resp, err := c.Execute(ctx, req)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp.Status)
}
```

------------------------------------------------------------------------

## Roadmap

Initial development tasks:

-   [ ] Core transport client interface
-   [ ] Request / response primitives
-   [ ] Authentication strategies
-   [ ] HTTP/API transport client
-   [ ] SFTP transport client
-   [ ] Retry and timeout helpers
-   [ ] Client registry
-   [ ] Transport middleware support
-   [ ] Webhook utilities

------------------------------------------------------------------------

## License

💡 Originally created by [Isidro A. Lopez
G.](https://github.com/ialopezg)\
🏢 Maintained by the [Entiqon Labs
Organization](https://github.com/entiqon)

[MIT](./LICENSE)
