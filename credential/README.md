# Credential Strategies

Package: `github.com/entiqon/transport/credential`

The `credential` package provides authentication credential strategies
for the transport library.

These credentials implement the `auth.Credential` interface and inject
authentication information into outgoing HTTP requests.

Credential strategies remain independent of the transport client and
can be applied through transport configuration options.

---

## Supported Credentials

### Access Token

The Access Token credential injects a token into a configurable HTTP
header.

This pattern is commonly used by APIs that rely on custom header tokens.

Examples include:

* Shopify
* internal service APIs
* partner integrations

The credential sets the header value directly on the outgoing request.

---

### Bearer Token

The Bearer Token credential injects the standard HTTP Authorization
header using the Bearer authentication scheme.

This pattern is widely used by APIs implementing OAuth2 or static bearer
authentication.

The credential sets the header:

Authorization: Bearer `<token>`

---

### API Key

The API Key credential injects a static key into outgoing HTTP requests.

The key can be applied either as:

* an HTTP header
* a query parameter

This pattern is commonly used by third-party APIs that rely on
simple API key authentication.

Example header usage:

X-API-Key: `<key>`

Example query usage:

https://api.example.com/resource?api_key=`<key>`

---

### Basic Authentication

The Basic credential injects an HTTP Authorization header using the
Basic authentication scheme.

The credential sets the header:

Authorization: Basic `<base64(username:password)>`

This pattern is commonly used by legacy APIs and service integrations.

---

## Design

Credential strategies follow the transport credential abstraction.

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

The `credential` package intentionally provides only simple credential
strategies.

It does not implement:

* OAuth2 authorization flows
* token refresh logic
* credential resolution
* session management

These responsibilities belong to higher-level components in the
consuming application.
