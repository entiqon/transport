# entiqon/transport

`transport` is a minimal Go library that provides reusable primitives for executing communication across different transport channels.

The library focuses strictly on the **communication layer**, allowing applications to interact with external systems through a unified transport abstraction.

It is designed to remain **small, composable, and transport-focused**, leaving business logic and orchestration to the application layer.

---

## Goals

- Provide a minimal transport client abstraction
- Support multiple communication channels (API, SFTP, Webhooks, etc.)
- Offer reusable authentication primitives
- Keep configuration models simple and transport-specific
- Remain lightweight and easy to integrate

---

## Intent

`transport` is not an integration framework.

It intentionally avoids:

- business workflows
- domain integrations
- orchestration logic
- data transformation

Those concerns belong to the application using the library.

---

## Example

A simple example showing how a transport client executes a request.

```go
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

	// In real usage the client would be an API, SFTP, or other transport implementation.
	var c client.Client

	resp, err := c.Execute(ctx, req)
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Status)
}
```

---

## Roadmap

Initial development tasks:

- [ ] Core transport client interface
- [ ] Request / response primitives
- [ ] Authentication strategies
- [ ] HTTP/API transport client
- [ ] SFTP transport client
- [ ] Retry and timeout helpers
- [ ] Client registry
- [ ] Transport middleware support
- [ ] Webhook utilities

---

## 📄 License

💡 Originally created by [Isidro A. Lopez G.](https://github.com/ialopezg)  
🏢 Maintained by the [Entiqon Labs Organization](https://github.com/entiqon)

[MIT](./LICENSE) — © Isidro A. López G. / Entiqon Project
