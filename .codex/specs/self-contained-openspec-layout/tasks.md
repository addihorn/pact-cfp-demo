# Implementation Plan

- [x] 1. Create self-contained OpenSpec layout proposal
  - Add requirements, design, and task documents under `.codex/specs/self-contained-openspec-layout/`.
  - Capture acceptance criteria for canonical `.codex/` usage and absence of parallel OpenSpec roots.
  - _Requirements: 1.1, 1.2, 1.3, 1.4, 2.1, 2.2, 2.3, 3.1, 3.2, 3.3_

- [x] 2. Update project index documentation
  - Add an `OpenSpec Layout` section to `project.md`.
  - State that `.codex/` is the canonical OpenSpec-compatible root.
  - State that the repository is self-contained and needs no external OpenSpec setup.
  - _Requirements: 1.1, 1.2, 1.3_

- [x] 3. Update agent guidance
  - Add `OpenSpec Compatibility` guidance to `AGENTS.md`.
  - Document canonical OpenSpec context files for agents.
  - Document future feature spec placement under `.codex/specs/<feature-name>/`.
  - _Requirements: 2.1, 2.2, 2.3, 3.1, 3.2, 3.3_

- [x] 4. Verify documentation consistency
  - Confirm no documentation now implies that `openspec/AGENTS.md` or top-level `specs/` are required.
  - Confirm the final layout matches the design document.
  - _Requirements: 1.3, 3.2, 3.3_
