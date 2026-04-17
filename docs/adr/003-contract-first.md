# ADR-003: Contract-First API

## Status
Accepted

## Decision
Use `docs/openapi.yaml` as the single source of truth for the Todo API contract.

## Consequences
- API documentation remains versionable and reviewable.
- Frontend request/response shapes can be validated against the same contract.