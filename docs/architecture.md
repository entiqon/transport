# Architecture

The transport library follows a layered communication architecture
designed to keep request execution independent from authentication
and application concerns.

Application
    ↓
Transport Client
    ↓
Credential Strategy
    ↓
External System

## Layers

### Application

The consuming application defines business logic, request creation,
and response handling.

The transport library intentionally does not include any domain logic
or workflow orchestration.

---

### Transport Client

The transport client is responsible for executing communication
requests against external systems.

Responsibilities:

- validate transport requests
- construct the underlying protocol request (e.g., HTTP)
- apply request headers and query parameters
- apply credentials when configured
- execute the request
- normalize the response

The primary transport implementation is provided in:

client/api

---

### Credential Strategy

Credential strategies modify outgoing requests to apply authentication
information required by external systems.

Examples include:

- Access Token headers
- Bearer tokens
- API keys
- Signed requests

Credential strategies implement the following interface:

```go
type Credential interface {
    Apply(ctx context.Context, req *http.Request) error
}
```

This design keeps authentication logic independent from the transport
client implementation.

---

### External System

The external system represents any HTTP-based service or API that the
transport client communicates with.

Examples:

- REST APIs
- partner integrations
- internal services

---

## Design Principles

### Separation of Concerns

Each layer has a clearly defined responsibility:

| Layer | Responsibility |
|------|----------------|
| Application | Business logic and workflows |
| Transport Client | Request execution |
| Credential Strategy | Request authentication |
| External System | Target API or service |

---

### Composability

Transport clients are configured using functional options.

Example:

```go
client := api.New(
    api.WithCredential(token.NewBearerToken("token")),
)
```

This allows flexible composition of transport behavior without
increasing complexity.

---

### Minimal Surface

transport intentionally avoids responsibilities outside the
communication layer.

The library does not include:

- business workflows
- domain logic
- orchestration
- data transformation
- response decoding

These concerns belong to the consuming application.

---

### Transport Independence

The Request and Response types abstract the underlying protocol
implementation.

This allows additional transports (for example gRPC, message queues,
or other protocols) to be introduced without changing the public
interface.
