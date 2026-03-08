# Changelog

All notable changes to this project will be documented in this file.

The project follows **Semantic Versioning**.

---

## [0.2.0] - 2026-03-08

### Added

- Authentication abstraction via `auth.Auth`
- `AccessToken` authentication strategy
- Authentication support in the API client through `WithAuth`

### Improved

- Expanded API client examples
- Improved unit test coverage for transport execution paths
- Clearer GoDoc documentation for transport primitives

### Changed

- Refactored authentication logic to fully decouple it from the transport client

### Removed

- Unused token provider abstraction

---

## [0.1.0] - Initial Release

### Added
- Initial transport library foundation
- HTTP API transport client (`client/api`)
- Authentication interface (`auth.Auth`)
- Request and Response primitives
- Functional client options (`WithHTTPClient`, `WithAuth`)
- Retry helper utility (`helpers.Retry`)
- Structured unit tests
- API client documentation
- Project README and GoDoc