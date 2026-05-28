# AGENTS.md Contract

This file defines the agent contract and serves as the index to project guidance. It applies repo-wide unless overridden by a nested AGENTS.md.

## Steering Documents
- Purpose: Treat Steering as the first-party source for product, technical, and structural guidance.
- Location: `.codex/steering/product.md`, `.codex/steering/tech.md`, `.codex/steering/structure.md`.
- Access: Use repository-provided helpers to resolve paths; avoid hardcoded absolutes.
- Mutations: Treat Steering as read-only; follow the established change-management flow instead of ad-hoc edits.
- Usage: Summarize applicable Steering points in outputs; cite the specific file/section without restating details.

## Decision Precedence
1) Steering documents under `.codex/steering/*.md`.
2) This AGENTS.md contract (general repository conventions).
3) Generic assumptions (avoid unless explicitly allowed).

## OpenSpec Compatibility
- This repository intentionally uses a Kiro/Codex-style OpenSpec layout under `.codex/`.
- Treat root `AGENTS.md`, `project.md`, `.codex/steering/*.md`, and `.codex/specs/<feature-name>/*.md` as the canonical OpenSpec context.
- The absence of `openspec/AGENTS.md`, top-level `specs/**/*.spec.md`, or `specs/<feature>/spec.md` is intentional and does not indicate missing setup.
- New feature specifications MUST be created under `.codex/specs/<feature-name>/`.
- New feature specifications SHOULD contain `requirements.md`, `design.md`, and `tasks.md`.
- Do not create parallel OpenSpec documents under `openspec/` or top-level `specs/` unless the project formally migrates layouts.
- If an OpenSpec CLI configuration file is present, it must act only as an adapter that points tools back to `.codex/` and the repository-owned context files.
- When generic OpenSpec directory assumptions conflict with this repository guidance, this repository guidance takes precedence.

## Agent Behavior Contract
- Prefer project-provided abstractions for CLI operations; avoid ad-hoc process spawning when wrappers exist.
- Respect feature flags and configuration gating documented in Steering or related code references.
- Logging: Use the shared logging utilities; record key lifecycle and error paths without excess noise.
- Error handling: Route failures through centralized services/utilities rather than ad-hoc try/catch loops.
- Performance/UX: Honor project guidelines for long-running work so the host environment remains responsive.
- Reference Steering for specifics (naming, boundaries, directory layout) rather than restating them here.

## Paths and I/O
- Workspace I/O: Prefer project-sanctioned file system helpers for read/write/create operations.
- Path resolution: Use repository-approved path utilities for `.codex/steering` and related directories; avoid absolute paths.
- Write boundaries: Only modify files within approved workspace areas defined by project rules.
- Steering: Do not overwrite files under `.codex/steering` directly; rely on the sanctioned update flow.

## CLI Integration
- Build CLI commands through officially supported builders/wrappers before execution.
- Approval modes and model flags: Reference their definition sites/tests without duplicating values.
- Verify required tooling availability before invocation; surface setup guidance if unavailable.

## Submission Checklist (For Agents)
- Verify decisions against `.codex/steering/*.md`; cite files/sections without duplication.
- Resolve steering paths via approved utilities; avoid absolute paths.
- Respect feature flags and constraints documented by the project.
- Use project-sanctioned wrappers for CLI interactions.
- Avoid restating Steering or code constants; keep AGENTS.md concise and index-like.

## Non-Goals and Anti-Patterns
- Do not bypass official wrappers/utilities for CLI calls when they exist.
- Do not store state in globals beyond established singletons.
- Do not write outside approved directories or overwrite Steering directly.
- Do not re-enable disabled features unless explicitly requested.
