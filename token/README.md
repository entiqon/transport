# Token Credentials

Package: `github.com/entiqon/transport/token`

The `token` package provides token-based credential strategies for the
transport library.

These credentials implement the `auth.Credential` interface and inject
authentication information into outgoing HTTP requests.

Token credentials are independent from the transport client and can be
applied through transport configuration options.

---

## Supported Credentials

### Access Token

The Access Token credential injects a token into a configurable HTTP
header.

This pattern is commonly used by APIs that rely on custom header tokens.

Examples include:

- Shopify
- internal service APIs
- partner integrations

The credential sets the header value directly on the outgoing request.

---

### Bearer Token

The Bearer Token credential injects the standard HTTP Authorization
header using the Bearer authentication scheme.

This pattern is widely used by APIs implementing OAuth2 or static bearer
authentication.

The credential sets the header:

Authorization: Bearer <token>

---

## Design

Token credentials follow the transport credential abstraction.

Each credential modifies the outgoing HTTP request before it is executed
by the transport client.

Application
    ↓
Transport Client
    ↓
Credential Strategy
    ↓
External System

This design keeps authentication logic independent from request
execution.

---

## Package Scope

The `token` package intentionally provides only simple token-based
credentials.

It does not implement:

- OAuth2 authorization flows
- token refresh logic
- credential resolution
- session management

These responsibilities belong to higher-level components in the
consuming application.
