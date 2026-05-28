# Pact CFP Demo Project

## Overview

This repository is a 3-tier Todo demonstration application used to show service boundaries, REST contracts, local development workflow, and contract-oriented API behavior.

## Product Context

- Domain: single-tenant Todo CRUD with completion state and filtering.
- Audience: engineers learning service boundaries and stakeholders reviewing architecture decomposition.
- Current scope excludes authentication, authorization, multi-tenancy, and external integrations.

Primary source: `.codex/steering/product.md`.

## Technical Context

- Frontend: React + Vite + TypeScript in `apps/web`.
- Backend: Go + Chi + GORM in `apps/api`.
- Database: PostgreSQL for the composed environment, with local runtime database selection planned in existing specs.
- API style: REST.
- Canonical API contract: `docs/openapi.yaml`.

Primary source: `.codex/steering/tech.md`.

## Structure Context

- `apps/web`: UI tier and browser client code.
- `apps/api`: API tier, business logic, and persistence access.
- `infra`: operational notes and environment assets.
- `docs`: architecture, API, runbook, and ADRs.
- `.codex/specs`: specification-driven development documents.

Primary source: `.codex/steering/structure.md`.

## OpenSpec Layout

This project uses a Kiro/Codex-style OpenSpec layout. The canonical OpenSpec-compatible root is `.codex/`.

OpenSpec-relevant project context is stored in:

- `AGENTS.md`: repository agent contract and guidance index.
- `project.md`: concise project and OpenSpec layout index.
- `.codex/steering/product.md`: product intent, scope, audience, and non-goals.
- `.codex/steering/tech.md`: technical stack, contracts, runtime, and testing guidance.
- `.codex/steering/structure.md`: repository layout, boundaries, and documentation placement rules.
- `.codex/specs/<feature-name>/requirements.md`: feature requirements and acceptance criteria.
- `.codex/specs/<feature-name>/design.md`: architecture and design plan.
- `.codex/specs/<feature-name>/tasks.md`: implementation task breakdown.
- `openspec/config.yaml`, when present: OpenSpec CLI adapter config that points tools back to `.codex/` and repository-owned context files.

The repository is self-contained for OpenSpec-driven work. No external OpenSpec directory setup is required, and this project does not require `openspec/AGENTS.md`, OpenSpec documents under `openspec/`, or top-level `specs/` directories.

## Planning Rules

- Use OpenSpec to plan changes before implementation when the change affects behavior, architecture, contracts, or cross-tier boundaries.
- Keep implementation tasks traceable to requirements.
- Update `docs/openapi.yaml` as part of implementation planning whenever API request/response behavior changes.
- Capture durable architecture decisions in `docs/adr/` when appropriate.
