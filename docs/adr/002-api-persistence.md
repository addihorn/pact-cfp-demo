# ADR-002: API and Persistence Choices

## Status
Accepted

## Decision
Use REST + OpenAPI for service contract, Chi for routing, and GORM for PostgreSQL persistence.

## Consequences
- REST contract is easy to demo and inspect.
- GORM accelerates repository implementation while keeping SQL DB behavior explicit.