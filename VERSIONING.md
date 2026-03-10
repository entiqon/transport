# Versioning Strategy

This project follows **Semantic Versioning (SemVer)**.

Version format:

MAJOR.MINOR.PATCH

Example:

v0.8.0

---

# Pre-1.0 Releases

Before **v1.0.0**, the project is considered under active development.

During this phase:

- Breaking changes **may occur**
- Public APIs **may evolve**
- Internal architecture **may change**
- Canonical models **may still evolve**

Minor versions may include:

- new features
- refactors
- breaking changes

Patch versions typically include:

- bug fixes
- documentation improvements
- test updates

---

# Post-1.0 Releases

After **v1.0.0**, the public API is considered **stable**.

Versioning rules follow strict SemVer:

MAJOR  
Breaking API changes

MINOR  
Backward-compatible features and improvements

PATCH  
Bug fixes and documentation updates only

---

# Development Builds

Development builds may use the following tag format:

v0.x.y-dev.z

Example:

v0.9.0-dev.1

These builds are **not considered stable releases** and may change without notice.

---

# v1.0 Milestone

Version **v1.0.0** will be released once:

- Transport primitives are stable
- Credential strategies are finalized
- Provider architecture is stable
- Canonical models are considered stable