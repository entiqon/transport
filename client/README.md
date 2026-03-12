
# Transport Client

Directory: `transport/client`

The `client` module contains transport implementations responsible for
executing outbound communication with external systems.

The goal of this layer is to provide **transport executors** that can be
used by higher-level integrations while remaining independent from
business logic and authentication mechanisms.

Clients focus strictly on **communication protocols** (HTTP, SFTP, etc.)
and delegate authentication behavior to the `auth` layer.

---

# Package Structure

```
client
└── api
```

Currently the module provides an HTTP transport client, with additional
transport implementations planned as the library evolves.

---

# Packages

## api

Package:

```
transport/client/api
```

The `api` package implements an **HTTP transport client** built on top of
Go's `net/http` package.

This client is responsible for:

- constructing HTTP requests
- applying headers and query parameters
- applying authentication credentials
- executing the request using an HTTP client
- returning a normalized transport response

The API client does **not implement authentication logic directly**.
Authentication is applied using:

- credential strategies (`auth.Credential`)
- authentication providers (`auth.Provider`)

This keeps the HTTP transport reusable across different integrations.

---

# Core Components

### Client

Represents the transport executor.

```
Execute(ctx context.Context, req *Request) (*Response, error)
```

Responsibilities:

- validate transport request
- construct `http.Request`
- apply headers and query parameters
- apply authentication credentials
- execute the request
- return normalized response data

---

### Request

Represents an HTTP transport request.

Fields typically include:

- HTTP method
- endpoint or URL
- headers
- query parameters
- request body

The body is represented using `io.Reader` to support flexible streaming
of payloads.

---

### Response

Represents the normalized transport response.

Includes:

- HTTP status code
- response headers
- raw response body

The response body is returned as raw bytes so callers can perform custom
decoding depending on the integration.

---

# Design Principles

The client layer follows several design principles.

### Transport Isolation

Clients are responsible only for **communication mechanics**.

They do not contain:

- business logic
- integration workflows
- authentication flows

These concerns belong to other layers.

---

### Authentication Delegation

Authentication is handled externally through the `auth` package.

The client simply applies credentials that implement the
`auth.Credential` interface.

This allows integrations to plug in different authentication mechanisms
without modifying the client.

---

### Minimal Abstraction

The transport client remains intentionally thin and close to Go's
`net/http` behavior.

This approach ensures:

- predictable behavior
- easier debugging
- minimal abstraction overhead

---

# Roadmap

The `client` module will evolve as additional transport needs emerge.

Planned enhancements include the following.

---

## Retry Policies

Support for configurable retry strategies such as:

- exponential backoff
- retry on specific HTTP status codes
- retry on network failures

---

## Rate Limiting

Optional rate limiting controls for outbound requests.

Potential features:

- per-host limits
- burst capacity
- configurable throttling

---

## Middleware Pipeline

Composable middleware for request execution.

Examples:

- logging middleware
- tracing middleware
- metrics collection
- request mutation

---

## Observability Hooks

Support for transport-level instrumentation including:

- OpenTelemetry tracing
- request metrics
- structured logging

---

## Additional Clients

Future transport clients may include:

- `sftp` client
- `grpc` client
- `webhook` client
- streaming transports

These clients will follow the same philosophy:

- protocol-focused
- authentication delegated
- minimal abstraction

---

# License

© Entiqon Labs

Released under the MIT License.
