# Contributing

## Workflow
- Use feature branches.
- Keep changes scoped by tier (`apps/web`, `apps/api`, `infra`, `docs`).
- Update docs for behavior or interface changes.

## Quality Gates
- API: `go test ./...`, `go vet ./...`
- Web: `npm test -- --run`, `npm run lint`, `npm run build`
- Integration: validate `docker compose up --build`

## Architecture Rules
- Keep dependency direction: web -> api -> db.
- Do not let handler logic bypass service layer.
- Do not access DB directly from handlers.

## API Contract
- Treat `docs/openapi.yaml` as the canonical API contract.
- Keep API responses and frontend types aligned with the contract.