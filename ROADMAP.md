# Transport Authentication Roadmap

This roadmap outlines the authentication and client evolution planned
for the transport library.

The project separates **transport execution**, **credential
strategies**, and **credential providers** to keep the architecture
clean and extensible.

Transport execution must remain independent from authentication logic.

------------------------------------------------------------------------

# v0.2.0 --- Foundation

Authentication abstractions and initial credential strategy.

-   [x] `auth.Credential` strategy interface
-   [x] Authentication injection via `WithCredential`
-   [x] `AccessToken` credential strategy
-   [x] API client authentication integration
-   [x] Query parameter propagation
-   [x] Structured tests for credential execution

------------------------------------------------------------------------

# v0.3.0 --- Bearer Token

Standard bearer authentication used by most HTTP APIs.

-   [x] `BearerToken` credential strategy
-   [x] Authorization header injection
-   [x] Unit tests for bearer authentication
-   [x] API client documentation update

------------------------------------------------------------------------

# v0.4.0 --- API Key Authentication

Common authentication mechanisms used by most APIs.

-   [x] `APIKey` credential strategy
-   [x] API key injection via headers
-   [x] API key injection via query parameters
-   [x] Optional location (defaults to header)
-   [x] Location validation (`APIKeyLocation`)
-   [x] Expanded credential tests
-   [x] Minimal examples for pkg.go.dev

------------------------------------------------------------------------

# v0.5.0 --- Basic Authentication

HTTP Basic authentication strategy.

-   [x] `Basic` credential strategy
-   [x] Authorization header generation
    (`Basic base64(username:password)`)
-   [x] Unit tests for basic authentication
-   [x] Credential documentation updates
-   [x] Credential package normalization (`token` → `credential`)

------------------------------------------------------------------------

# v0.6.0 --- Request Signing (Initial Support)

Authentication mechanisms that require request hashing or signatures.

-   [x] `HMAC` request signing strategy

The advanced signing capabilities originally planned for this release
were postponed and will be delivered incrementally in **v0.9.x**.

------------------------------------------------------------------------

# v0.7.0 --- Credential Providers

Introduce dynamic credential resolution.

These components allow credentials to be obtained dynamically rather
than being statically configured.

-   [x] Credential resolver interface
-   [x] Static resolver
-   [x] Cached resolver
-   [x] Expiring token support
-   [x] Concurrency-safe token refresh

------------------------------------------------------------------------

# v0.8.0 --- OAuth2 Support

Dynamic authentication flows.

-   [x] OAuth2 client credentials resolver
-   [x] Refresh token resolver
-   [x] Automatic token refresh
-   [x] Token caching
-   [x] `WithAuthProvider` client option

------------------------------------------------------------------------

# v0.8.x --- API Client Stability

Stabilization of the API client behavior.

-   [x] Version header injection (`X-API-Version`)
-   [x] `WithVersion` client option
-   [x] `WithBasePath` support
-   [x] Base path normalization
-   [x] Safe URL resolution
-   [x] Query parameter propagation
-   [x] Header propagation
-   [x] Context-aware request execution
-   [x] Credential application pipeline
-   [x] Improved unit test coverage

------------------------------------------------------------------------

# v0.9.0 --- Request Body and Helpers

Introduce a request body abstraction and developer-friendly helpers for
common HTTP workflows.

-   [x] `Body` interface for typed request payloads
-   [x] `transport.JSON()` request body helper
-   [x] Retry-safe body readers
-   [ ] JSON response decoding helpers
-   [ ] `DoJSON` execution helper
-   [ ] Automatic `Content-Type` management
-   [ ] Response decoding utilities

Example:

    client.DoJSON(ctx, req, &result)

------------------------------------------------------------------------

# v0.9.x --- Advanced Request Signing

Completion of the request signing capabilities originally planned for
**v0.6.0**.

-   [ ] Configurable hashing algorithms
-   [ ] Payload signing support
-   [ ] Timestamp validation support
-   [ ] Canonical request builder

------------------------------------------------------------------------

# v0.10.0 --- Middleware Support

Introduce a middleware pipeline for request and response processing.

Middleware enables extensibility without modifying the transport client.

-   [ ] Middleware interface
-   [ ] Middleware execution pipeline
-   [ ] Logging middleware example
-   [ ] Metrics middleware example
-   [ ] Retry middleware example

Architecture:

    Request
       ↓
    Middleware chain
       ↓
    Credential strategy
       ↓
    HTTP execution

------------------------------------------------------------------------

# v0.11.0 --- Retry Policies

Add retry capabilities for transient failures.

-   [ ] Retry policy configuration
-   [ ] Exponential backoff support
-   [ ] Retry on network errors
-   [ ] Retry on configurable status codes
-   [ ] Retry middleware integration

Example:

    api.WithRetry(3)

------------------------------------------------------------------------

# v0.12.0 --- Observability

Add observability hooks to the transport client.

-   [ ] Structured logging support
-   [ ] Request duration tracking
-   [ ] Pluggable logger interface
-   [ ] OpenTelemetry integration hooks
-   [ ] Request tracing support

------------------------------------------------------------------------

# v0.13.0 --- Rate Limiting

Client-side rate limiting support.

-   [ ] Token bucket rate limiter
-   [ ] Configurable request limits
-   [ ] Burst configuration
-   [ ] Integration with middleware pipeline

Example:

    api.WithRateLimit(10, time.Second)

------------------------------------------------------------------------

# v1.0.0 --- Enterprise Authentication

Advanced authentication mechanisms.

-   [ ] AWS Signature V4 strategy
-   [ ] mTLS support
-   [ ] Pluggable credential providers
-   [ ] Advanced retry policies for authentication failures

------------------------------------------------------------------------

# Design Principles

The transport layer will always remain independent from authentication
logic.

    Client
       ↓
    Credential Strategy
       ↓
    Credential Resolver

This separation ensures the library can support static tokens, OAuth
flows, signed requests, and enterprise authentication mechanisms without
modifying the transport client.

Credential strategies remain simple request modifiers, while credential
resolvers handle dynamic token acquisition and lifecycle management.
