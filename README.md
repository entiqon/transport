# entiqon/transport

[![Go Reference](https://pkg.go.dev/badge/github.com/entiqon/transport.svg)](https://pkg.go.dev/github.com/entiqon/transport)
[![CI](https://github.com/entiqon/transport/actions/workflows/ci.yml/badge.svg)](https://github.com/entiqon/transport/actions/workflows/ci.yml)
[![codecov](https://codecov.io/gh/entiqon/transport/branch/main/graph/badge.svg)](https://codecov.io/gh/entiqon/transport)
[![Go Report Card](https://goreportcard.com/badge/github.com/entiqon/transport)](https://goreportcard.com/report/github.com/entiqon/transport)
[![Latest Release](https://img.shields.io/github/v/release/entiqon/transport)](https://github.com/entiqon/transport/releases)
[![License](https://img.shields.io/github/license/entiqon/transport)](https://github.com/entiqon/transport/blob/main/LICENSE)
[![Docs](https://img.shields.io/badge/docs-api--client-blue)](docs/api-client.md)

`transport` is a minimal Go library that provides reusable primitives
for executing requests across different communication transports.

The library focuses strictly on the **communication layer**, allowing
applications to interact with external systems through a unified
transport abstraction. It is designed to remain **small, composable, and
transport-focused**, leaving business logic, orchestration, and data
transformation to the application layer.

---

## Why transport?

Applications often interact with external systems through different
communication channels such as APIs, SFTP servers, or webhooks.

`transport` provides a small set of primitives that make those
interactions consistent while keeping application logic independent
of the underlying communication mechanism.

---

## Example

```go
package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/entiqon/transport/client/api"
)

func main() {

	ctx := context.Background()

	client := api.New(
		api.WithHTTPClient(http.DefaultClient),
	)

	req := &api.Request{
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

## Documentation

Detailed documentation is available in the `/docs` directory.

- [APIClient](docs/api-client.md) – Client design and usage
- Architecture – transport architecture overview

---

## Roadmap

- [x] Core transport client interface
- [x] Request / response primitives
- [x] HTTP/API transport client
- [x] Authentication abstraction
- [ ] SFTP transport client
- [ ] Retry and timeout helpers
- [ ] Client registry
- [ ] Transport middleware support
- [ ] Webhook utilities

---

## License

💡 Originally created by [Isidro A. Lopez G.](https://github.com/ialopezg)  
🏢 Maintained by the [Entiqon Labs](https://github.com/entiqon)

[MIT](./LICENSE) License
