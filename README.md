# 3-Tier Todo Demo

A single-repo Todo application demonstrating a 3-tier architecture:
- Web tier: React + Vite + TypeScript (`apps/web`)
- Service tier: Go + Chi + GORM (`apps/api`)
- Data tier: PostgreSQL (`docker-compose.yml`)

## Quickstart

Prerequisites:
- Docker + Docker Compose
- Go toolchain
- Node.js + npm

Run with Docker Compose:
```bash
docker compose up --build
```

Endpoints:
- Web: `http://localhost:5173`
- API: `http://localhost:8080`
- Health: `http://localhost:8080/healthz`
- Readiness: `http://localhost:8080/readyz`
- Metrics: `http://localhost:8080/metrics`

## Local Development

API:
```bash
cd apps/api
go run ./cmd/api
```

Web:
```bash
cd apps/web
npm install
npm run dev
```

## Database

Migrations are in `apps/api/migrations` and are executed by the `migrate` service in Compose.
Seed SQL is in `apps/api/seed/seed.sql`.

## Commands

```bash
make up
make down
make migrate
make seed
make test
make lint
```

## Documentation

- Architecture: `docs/architecture.md`
- API and OpenAPI: `docs/api.md`, `docs/openapi.yaml`
- Runbook: `docs/runbook.md`
- ADRs: `docs/adr/`
- Contribution guide: `CONTRIBUTING.md`