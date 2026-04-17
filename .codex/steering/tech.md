# Technical Steering

## Core Stack
- Frontend: React + Vite + TypeScript
- Backend: Go + Chi + GORM
- Database: PostgreSQL

## Contract and Interfaces
- Use REST for service interfaces.
- Maintain OpenAPI as the canonical API contract (`docs/openapi.yaml`).
- Keep frontend models aligned with API response fields.

## Runtime and Operations
- Use Docker Compose for local orchestration.
- Expose `/healthz`, `/readyz`, and `/metrics` from backend.
- Use SQL migrations and deterministic seed data for bootstrapping.

## Testing
- Backend: service-level unit tests and HTTP integration tests.
- Frontend: component smoke tests with Vitest and Testing Library.