# Design Document

## Overview

This design introduces runtime database backend selection for the API, defaulting to SQLite for local/non-production execution and using PostgreSQL for production. The implementation keeps the service and repository contracts unchanged by isolating backend selection in `internal/config` and `internal/db` bootstrap layers.

Steering alignment:
- Product: `.codex/steering/product.md` (`Product Scope (v1)` local-first workflow intent)
- Technical: `.codex/steering/tech.md` (`Core Stack`, `Testing`)
- Structure: `.codex/steering/structure.md` (`Boundary Rules`)

## Architecture

The current startup path calls a single Postgres connector. The updated architecture adds a selector that derives backend choice from `APP_ENV` and returns a `*gorm.DB` from the selected connector implementation.

```mermaid
flowchart TD
    A[cmd/api/main.go] --> B[config.Load()]
    B --> C{APP_ENV == production?}
    C -- yes --> D[db.PostgresConnector.Connect]
    C -- no --> E[db.SQLiteConnector.Connect]
    D --> F[*gorm.DB]
    E --> F[*gorm.DB]
    F --> G[repository.NewTodoRepository]
    G --> H[service.NewTodoService]
    H --> I[api.NewRouter]
```

Design decisions:
- Keep repository/service signatures unchanged (`*gorm.DB` remains the persistence handle).
- Use connector abstraction in `internal/db` to avoid backend conditionals outside bootstrap.
- Keep backend selection deterministic: `APP_ENV=production` => Postgres, all other values => SQLite.

## Components and Interfaces

### Config updates (`internal/config`)

Extend `Config` with DB selection inputs:
- `DatabaseDSN string` (existing; production use)
- `SQLitePath string` (new; default `apps/api/data/todos.db`)
- `Environment string` (existing `APP_ENV`, authoritative selector)

Behavior:
- `Environment == "production"` => backend `postgres`.
- Otherwise => backend `sqlite`.
- `SQLitePath` supports override via `SQLITE_PATH` env var.

### DB connector abstraction (`internal/db`)

Introduce a connector interface and selector:

```go
type Connector interface {
	Connect(cfg config.Config) (*gorm.DB, error)
	BackendName() string
}
```

Planned implementations:
- `PostgresConnector` using `gorm.io/driver/postgres`.
- `SQLiteConnector` using `gorm.io/driver/sqlite`.

Selector entrypoint:
- `ConnectFromConfig(cfg config.Config) (*gorm.DB, error)`
- Chooses connector from `cfg.Environment`.
- Wraps errors with backend context for actionable startup failures.

### Startup wiring (`cmd/api/main.go`)

- Replace direct `db.Connect(cfg.DatabaseDSN)` call with `db.ConnectFromConfig(cfg)`.
- Keep downstream wiring unchanged (`repository.NewTodoRepository(database)` etc.).

### Dependency update (`apps/api/go.mod`)

- Add `gorm.io/driver/sqlite`.

## Data Models

No API contract or domain model changes are required.

Persistence-level considerations:
- Existing GORM `domain.Todo` model remains unchanged.
- No repository method changes.
- SQLite database file is created at configured path during connect/open.
- Migration strategy for SQLite in this feature scope:
  - Use GORM auto-migration for local SQLite mode to ensure schema availability without Postgres migration tooling changes.
  - Keep existing SQL migrations unchanged for PostgreSQL path.

Rationale:
- Avoids coupling local SQLite startup to Postgres-specific migration syntax (`BIGSERIAL`, `TIMESTAMPTZ`).
- Keeps local debug workflow fast and deterministic.

## Error Handling

Startup errors must fail fast with context:
- Production mode + missing/invalid DSN => explicit configuration/startup error.
- SQLite mode + invalid/unwritable file path => startup error including path and backend name.
- Selector must not silently fall back from production to SQLite.

Error format guidance:
- Include backend name (`postgres` or `sqlite`) and root cause (`%w` wrapping).
- Keep startup log line concise and actionable.

## Testing Strategy

### Unit tests: config and selection

- `internal/config/config_test.go`
  - `APP_ENV=production` selects Postgres mode.
  - unset `APP_ENV` selects SQLite mode.
  - non-production values (`local`, `dev`, `test`) select SQLite mode.
  - default and override `SQLITE_PATH` behavior.

- `internal/db/connector_test.go`
  - `ConnectFromConfig` routes to Postgres in production.
  - `ConnectFromConfig` routes to SQLite otherwise.
  - Production misconfig and SQLite path failures return backend-specific wrapped errors.

### Regression tests

- Run existing API/service tests to confirm no interface churn:
  - `internal/service/todo_service_test.go`
  - `internal/api/router_integration_test.go`

### Optional lightweight integration

- Add an in-process SQLite connection test for basic CRUD repository behavior to validate local mode end-to-end in code without Docker dependencies.
