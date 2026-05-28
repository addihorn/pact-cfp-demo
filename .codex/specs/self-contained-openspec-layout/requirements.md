# Requirements Document

## Introduction

This feature formalizes the repository's Kiro/Codex-style OpenSpec layout so OpenSpec-driven work can run from the checked-in repository without creating parallel `openspec/` or top-level `specs/` directories. The project must be self-contained and store its agent guidance, steering context, and feature specifications in canonical in-repository files.

## Requirements

### Requirement 1

**User Story:** As an agent or contributor, I want the repository to define its OpenSpec-compatible layout explicitly, so that I can discover project guidance without external setup.

#### Acceptance Criteria

1. WHEN project documentation describes OpenSpec usage THEN it SHALL identify `.codex/` as the canonical OpenSpec-compatible root.
2. WHEN OpenSpec-relevant files are listed THEN the list SHALL include root `AGENTS.md`, `project.md`, `.codex/steering/*.md`, and `.codex/specs/<feature-name>/*.md`.
3. WHEN a contributor starts OpenSpec-driven planning THEN they SHALL NOT need to create `openspec/AGENTS.md`, top-level `specs/`, or other external scaffolding.
4. IF an OpenSpec CLI configuration file exists THEN it SHALL point tools back to `.codex/` and repository-owned context files instead of defining a parallel spec root.

### Requirement 2

**User Story:** As a maintainer, I want future specs to follow one layout, so that planning artifacts do not diverge across directories.

#### Acceptance Criteria

1. WHEN a new feature specification is created THEN it SHALL be placed under `.codex/specs/<feature-name>/`.
2. WHEN a new feature specification is created THEN it SHOULD include `requirements.md`, `design.md`, and `tasks.md`.
3. WHEN contributors look for active or historical specs THEN they SHALL treat `.codex/specs/` as the source of truth.

### Requirement 3

**User Story:** As an agent, I want root guidance to explain the repository-specific OpenSpec convention, so that generic OpenSpec expectations do not override this project's chosen layout.

#### Acceptance Criteria

1. WHEN an agent reads `AGENTS.md` THEN it SHALL find explicit OpenSpec compatibility guidance for the `.codex/` layout.
2. WHEN `openspec/AGENTS.md` or top-level `specs/` directories are absent THEN agents SHALL treat that absence as intentional.
3. WHEN there is a conflict between generic OpenSpec directory assumptions and repository guidance THEN the repository guidance SHALL take precedence.
