# Changelog

All notable changes to this project will be documented in this file.

The project follows **Semantic Versioning**.

---

# [0.7.0] - 2026-03-09

## Added

* `HMAC` credential strategy for request signing
* HMAC-SHA256 signature generation for outgoing HTTP requests
* Automatic injection of request signing headers:
    * `X-Key`
    * `X-Timestamp`
    * `X-Signature`
* Unit tests covering HMAC credential validation and execution

## Improved

* Expanded credential documentation
* README examples updated to include HMAC request signing
* Credential strategy coverage extended across the library

## Compatibility

Fully backward compatible with **v0.6.0**.

---

# [0.6.0] - 2026-03-09

## Added

* `JWT` credential strategy
* Automatic Bearer scheme when using `Authorization` header
* Support for custom JWT headers (e.g. `X-JWT-Assertion`)
* Table-driven credential tests covering all strategies

## Improved

* Expanded credential documentation
* Improved pkg.go.dev examples
* Consistent credential strategy coverage across tests

## Compatibility

Fully backward compatible with **v0.5.0**.

---

# [0.5.0] - 2026-03-09

## Added

* `Basic` credential strategy implementing HTTP Basic authentication
* Authorization header generation using `Basic base64(username:password)`
* Unit tests for Basic authentication
* Credential documentation updates reflecting new strategy

## Changed

* Renamed `token` package to `credential`
* Updated public API to reflect credential-based naming:

    * `credential.AccessToken`
    * `credential.BearerToken`
    * `credential.APIKey`
    * `credential.Basic`
* Updated package documentation and README to reflect credential terminology
* Updated GoDoc comments and examples across the repository

## Improved

* Normalized credential strategy naming across the library
* Improved documentation consistency between README, GoDoc, and examples

## Compatibility

This release is backward compatible at the transport level, but the
`token` package has been replaced by `credential`. Applications should
update imports accordingly.

---

# [0.4.0] - 2026-03-09

## Added

* `APIKey` credential strategy supporting:
    * header injection
    * query parameter injection
* Optional API key location (defaults to **header** when not provided)
* Validation for API key configuration:
    * missing key name
    * missing value
    * invalid location
* Minimal usage examples for token credentials on **pkg.go.dev**
* Expanded credential tests for APIKey usage

## Improved

* Expanded credential documentation across **README** and **GoDoc**
* Improved examples for:

    * `AccessToken`
    * `BearerToken`
    * `APIKey`
* Improved test coverage across authentication strategies
* Minor documentation refinements across the repository

## Compatibility

This release is fully backward compatible with **v0.3.0**.

---

# [0.3.0] - 2026-03-09

## Added

* `BearerToken` credential strategy
* Credential injection in the API client via `WithCredential`
* Header helper methods for `Response`
* Improved transport execution tests

## Changed

* Replaced authentication abstraction with credential-based model
* `WithAuth` option replaced by `WithCredential`
* Updated API client documentation and examples
* Updated architecture documentation to reflect credential strategies

## Improved

* Expanded unit test coverage for credentials and execution paths
* Improved GoDoc for transport primitives and credential interfaces
* Updated repository documentation (README, architecture, API docs)

---

# [0.2.0] - 2026-03-08

## Added

* Authentication abstraction via `auth.Auth`
* `AccessToken` authentication strategy
* Authentication support in the API client through `WithAuth`

## Improved

* Expanded API client examples
* Improved unit test coverage for transport execution paths
* Clearer GoDoc documentation for transport primitives

## Changed

* Refactored authentication logic to fully decouple it from the transport client

## Removed

* Unused token provider abstraction

---

# [0.1.0] - Initial Release

## Added

* Initial transport library foundation
* HTTP API transport client (`client/api`)
* Authentication interface (`auth.Auth`)
* Request and Response primitives
* Functional client options (`WithHTTPClient`, `WithAuth`)
* Retry helper utility (`helpers.Retry`)
* Structured unit tests
* API client documentation
* Project README and GoDoc
