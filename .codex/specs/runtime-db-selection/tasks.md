# Implementation Plan

- [x] 1. Establish database connector abstraction and shared contracts
  - Define a connector interface in API infrastructure that initializes and returns the same GORM DB type already used by repositories.
  - Add config parsing helpers for `APP_ENV`, PostgreSQL DSN, and SQLite path values.
  - Add unit tests validating config parsing and connector selection inputs.
  - _Requirements: 3.1, 3.2, 4.1, 4.2_

- [x] 2. Implement PostgreSQL connector for production mode
  - Create a PostgreSQL connector implementation that opens a GORM PostgreSQL connection from configured DSN.
  - Return clear startup errors for missing/invalid DSN with backend context.
  - Add tests for successful production connector initialization and configuration failure paths.
  - _Requirements: 2.1, 2.2, 2.3, 4.4_

- [ ] 3. Implement SQLite connector for local-first mode
  - Create a SQLite connector implementation that opens a file-based GORM SQLite connection.
  - Implement default file path fallback to `apps/api/data/todos.db` when no SQLite path is configured.
  - Add tests for default path behavior, explicit path override behavior, and initialization failure handling.
  - _Requirements: 1.1, 1.2, 1.3, 1.4, 1.5, 4.3, 4.4_

- [ ] 4. Add environment-based connector selection in startup wiring
  - Implement selection logic so `APP_ENV=production` chooses PostgreSQL and all other/unset values choose SQLite.
  - Wire selection in the API bootstrap/composition root only, keeping repository and service usage unchanged.
  - Add startup-level tests covering production selection, non-production selection, and no-fallback behavior in production.
  - _Requirements: 1.1, 1.2, 1.3, 2.1, 2.3, 3.1, 3.3, 4.1, 4.2_

- [ ] 5. Refactor existing DB initialization to use connector abstraction
  - Replace direct database initialization branches with the connector factory/selection path.
  - Preserve existing repository/service constructor signatures and data-access call sites.
  - Add regression tests to ensure CRUD flows behave identically without backend-specific branching in business logic.
  - _Requirements: 3.2, 3.3, 3.4, 4.5_

- [ ] 6. Strengthen startup and connector error reporting
  - Standardize startup error wrapping to include selected backend and root cause.
  - Ensure errors are returned early from bootstrap when connector initialization fails.
  - Add tests asserting actionable error messages for PostgreSQL and SQLite failure scenarios.
  - _Requirements: 1.5, 2.2, 4.4_

- [ ] 7. Validate end-to-end automated test coverage across modes
  - Add or update automated integration tests to run representative API flows against SQLite mode.
  - Verify existing service/API tests continue passing without service/repository interface changes.
  - Add a production-mode startup test harness that confirms PostgreSQL selection behavior without SQLite fallback.
  - _Requirements: 2.3, 3.3, 4.1, 4.2, 4.5_
