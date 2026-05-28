# Design Document

## Overview

This design makes the existing `.codex/` documentation layout the explicit OpenSpec-compatible convention for the project. It avoids duplicate planning roots and keeps all required context inside the repository.

Steering alignment:
- Product: `.codex/steering/product.md` (`Purpose`, `Product Scope (v1)`)
- Technical: `.codex/steering/tech.md` (`Contract and Interfaces`, `Testing`)
- Structure: `.codex/steering/structure.md` (`Repository Layout`, `Documentation Rules`)

## Architecture

The repository will use this canonical documentation graph:

```text
AGENTS.md
project.md
.codex/
  steering/
    product.md
    tech.md
    structure.md
  specs/
    <feature-name>/
      requirements.md
      design.md
      tasks.md
```

Design decisions:
- Keep `AGENTS.md` as the agent entry point.
- Keep `project.md` as the concise project/OpenSpec index.
- Keep `.codex/steering/` as the first-party product, technical, and structural guidance source.
- Keep `.codex/specs/` as the only feature-specification root.
- Do not introduce OpenSpec documents under `openspec/` or top-level `specs/` directories.
- Allow OpenSpec CLI configuration only as an adapter that points tools back to `.codex/` and repository-owned context files.

## Documentation Updates

### `project.md`

Add an `OpenSpec Layout` section that states:
- The project uses a Kiro/Codex-style OpenSpec layout.
- `.codex/` is the canonical OpenSpec-compatible root.
- The repository is self-contained for OpenSpec-driven work.
- No external OpenSpec setup is required.
- OpenSpec CLI configuration, when present, is an adapter rather than a spec root.

### `AGENTS.md`

Add `OpenSpec Compatibility` guidance that states:
- Agents should use root `AGENTS.md`, `project.md`, `.codex/steering/*.md`, and `.codex/specs/<feature-name>/*.md` as canonical OpenSpec context.
- The absence of `openspec/AGENTS.md` and top-level `specs/` is intentional.
- New specs belong under `.codex/specs/<feature-name>/`.
- OpenSpec CLI configuration must not define a competing documentation root.

## Data and Migration

No runtime data migration is required. This is a documentation and planning-layout change only.

## Error Handling

The documentation should prevent these failure modes:
- Agents reporting missing OpenSpec setup because `openspec/AGENTS.md` is absent.
- Contributors creating duplicate specs under top-level `specs/`.
- Feature planning becoming split between `.codex/specs/` and another OpenSpec root.
- OpenSpec CLI configuration drifting away from the canonical `.codex/` context.

## Testing Strategy

Manual verification:
- Confirm `project.md` documents `.codex/` as canonical.
- Confirm `AGENTS.md` documents OpenSpec compatibility and future spec placement.
- Confirm this proposal exists under `.codex/specs/self-contained-openspec-layout/` with `requirements.md`, `design.md`, and `tasks.md`.
