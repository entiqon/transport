# Architecture

transport follows a layered communication architecture.

Application
    ↓
Transport Client
    ↓
Authentication
    ↓
Token Provider
    ↓
External System

## Design Principles

### Separation of Concerns

Each layer has a single responsibility:

- Transport executes requests
- Auth modifies requests
- Token providers resolve credentials

### Composability

Transport clients are configured using functional options,
allowing flexible composition without increasing complexity.

### Minimal Surface

transport intentionally avoids:

- business workflows
- domain logic
- orchestration
- data transformation

These concerns belong to the consuming application.
