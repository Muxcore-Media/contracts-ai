# Changelog

## [0.1.0] - 2026-10-05

### Changed
- CI runs on GitHub-hosted runners from the umbrella template; retired-origin workflows removed.
- Dependencies resolve from published GitHub tags (no filesystem `replace`); requires core v0.6.0.

## [0.1.0] — 2026-09-08

### Added
- AI domain event constants and payloads in `events/`.
- Shared inference request/response types in `infer/`.
- `All()` and `Known()` registry helpers.
