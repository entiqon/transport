# Transport Authentication Roadmap

This roadmap outlines the authentication strategies planned for the
transport library.

Each release increment introduces new credential strategies while
keeping the transport layer independent of credential resolution.

---

# v0.2.0 — Foundation

Authentication abstractions and initial credential strategy.

* [x] `auth.Credential` strategy interface
* [x] Authentication injection via `WithCredential`
* [x] `AccessToken` credential strategy
* [x] API client authentication integration
* [x] Query parameter propagation
* [x] Structured tests for credential execution

---

# v0.3.0 — Bearer Token

Standard bearer authentication used by most HTTP APIs.

* [x] `BearerToken` credential strategy
* [x] Authorization header injection
* [x] Unit tests for bearer authentication
* [x] API client documentation update

---

# v0.4.0 — API Key Authentication

Common authentication mechanisms used by most APIs.

* [x] `APIKey` credential strategy
* [x] API key injection via headers
* [x] API key injection via query parameters
* [x] Optional location (defaults to header)
* [x] Location validation (`APIKeyLocation`)
* [x] Expanded credential tests
* [x] Minimal examples for pkg.go.dev

---

# v0.5.0 — Basic Authentication

HTTP Basic authentication strategy.

* [x] `Basic` credential strategy
* [x] Authorization header generation (`Basic base64(username:password)`)
* [x] Unit tests for basic authentication
* [x] Credential documentation updates
* [x] Credential package normalization (`token` → `credential`)

---

# v0.6.0 — Request Signing

Authentication mechanisms that require request hashing or signatures.

* [ ] `HMAC` request signing strategy
* [ ] Configurable hashing algorithms
* [ ] Payload signing support
* [ ] Timestamp validation support
* [ ] Canonical request builder

---

# v0.7.0 — Credential Resolvers

Introduce dynamic credential resolution.

These components allow credentials to be obtained dynamically rather
than being statically configured.

* [ ] Credential resolver interface
* [ ] Static resolver
* [ ] Cached resolver
* [ ] Expiring token support
* [ ] Concurrency-safe token refresh

---

# v0.8.0 — OAuth2 Support

Dynamic authentication flows.

* [ ] OAuth2 client credentials resolver
* [ ] Refresh token resolver
* [ ] Automatic token refresh
* [ ] Token caching

---

# v1.0.0 — Enterprise Authentication

Advanced authentication mechanisms.

* [ ] AWS Signature V4 strategy
* [ ] mTLS support
* [ ] Pluggable credential providers
* [ ] Advanced retry policies for authentication failures

---

# Design Principles

The transport layer will always remain independent from authentication
logic.

```
Client
   ↓
Credential Strategy
   ↓
Credential Resolver
```

This separation ensures the library can support static tokens, OAuth
flows, signed requests, and enterprise authentication mechanisms
without modifying the transport client.

Credential strategies remain simple request modifiers, while credential
resolvers handle dynamic token acquisition and lifecycle management.
