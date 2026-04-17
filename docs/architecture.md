# Architecture

## Tier Boundaries
- Web tier (`apps/web`): renders UI, calls backend only through HTTP.
- Service tier (`apps/api`): owns business rules and API contract.
- Data tier (PostgreSQL): owns persistence and query execution.

## Request Flow
1. Web sends REST request to API.
2. API handler validates payload and maps request to service input.
3. Service executes business rules and calls repository.
4. Repository runs DB operations via GORM.
5. API returns JSON response; web updates UI state.

## Data Flow
- Todo schema is defined in SQL migrations.
- API response shape is defined by `docs/openapi.yaml`.
- Frontend TypeScript models align to API response fields.

## Failure Boundaries
- Web surfaces API errors as user-facing messages.
- API maps domain errors to HTTP status codes.
- Readiness and health endpoints support operational checks.
- Metrics endpoint supports runtime observability.