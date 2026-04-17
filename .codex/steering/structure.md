# Structure Steering

## Repository Layout
- `apps/web`: UI tier and browser client code
- `apps/api`: API tier, business logic, and persistence access
- `infra`: operational notes and environment-level assets
- `docs`: architecture, API, runbook, and ADRs

## Boundary Rules
- Keep tier dependency direction as `web -> api -> db`.
- Keep handlers thin: validate/map input, delegate business rules to services.
- Keep repositories focused on data access concerns only.

## Documentation Rules
- Keep AGENTS.md reference-first and steering-aligned.
- Place architecture decisions in ADR files under `docs/adr`.