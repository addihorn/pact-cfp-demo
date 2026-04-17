# Requirements Document

## Introduction

This feature adds a local-first database runtime mode for the API so developers can start and debug backend code without requiring a running PostgreSQL instance. The API must choose the database backend from environment configuration, using SQLite by default for non-production environments and PostgreSQL for production. The database access behavior in services and repositories must remain stable, with backend selection isolated to bootstrap/infrastructure code.

## Requirements

### Requirement 1

**User Story:** As a backend developer, I want the API to default to a local SQLite database in non-production environments, so that I can run and debug the backend quickly without external infrastructure.

#### Acceptance Criteria

1. WHEN the API starts AND `APP_ENV` is unset THEN the system SHALL initialize a SQLite connection.
2. WHEN the API starts AND `APP_ENV` is `local` THEN the system SHALL initialize a SQLite connection.
3. WHEN the API starts AND `APP_ENV` is any non-`production` value THEN the system SHALL initialize a SQLite connection.
4. WHEN the API starts in SQLite mode AND no SQLite path is provided THEN the system SHALL use `apps/api/data/todos.db` as the default database file path.
5. IF SQLite initialization fails THEN the system SHALL return a startup error that includes the selected backend and failure reason.

### Requirement 2

**User Story:** As an operator, I want the API to use PostgreSQL in production, so that production data remains on the managed database service.

#### Acceptance Criteria

1. WHEN the API starts AND `APP_ENV` is `production` THEN the system SHALL initialize a PostgreSQL connection using the configured DSN.
2. IF `APP_ENV` is `production` AND the PostgreSQL DSN is missing or invalid THEN the system SHALL fail startup with a clear configuration error.
3. WHEN the API starts in production mode THEN the system SHALL NOT fall back to SQLite.

### Requirement 3

**User Story:** As a maintainer, I want database backend selection encapsulated behind infrastructure interfaces, so that business and repository logic remains backend-agnostic.

#### Acceptance Criteria

1. WHEN database selection is performed THEN the system SHALL use a database connector abstraction that determines the backend from configuration.
2. WHEN the backend connector is selected THEN the system SHALL expose the resulting database handle through the same GORM-based type currently consumed by repositories.
3. WHEN repositories and services execute CRUD behavior THEN the system SHALL require no backend-specific branching in repository/service methods.
4. IF a new backend is added later THEN the system SHALL allow adding a new connector implementation without changing service-layer interfaces.

### Requirement 4

**User Story:** As a developer, I want automated tests for environment-based backend selection and connector behavior, so that regressions are detected early.

#### Acceptance Criteria

1. WHEN automated config-selection tests run THEN they SHALL verify `APP_ENV=production` selects PostgreSQL.
2. WHEN automated config-selection tests run THEN they SHALL verify unset and non-production `APP_ENV` values select SQLite.
3. WHEN automated config-selection tests run THEN they SHALL verify default and override SQLite path behavior.
4. WHEN connector tests run THEN they SHALL verify failures are surfaced with actionable backend-specific errors.
5. WHEN existing service and HTTP integration tests run THEN they SHALL continue to pass without repository/service interface changes.
