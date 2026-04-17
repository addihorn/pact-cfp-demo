# Runbook

## Start Stack
```bash
docker compose up --build
```

## Stop Stack
```bash
docker compose down -v
```

## Apply Migrations
```bash
docker compose run --rm migrate
```

## Seed Data
```bash
docker compose run --rm seed
```

## Troubleshooting
- API cannot connect to DB: verify `postgres` health and `DATABASE_DSN`.
- Empty UI list: call `docker compose run --rm seed`.
- API not ready: check `GET /readyz` and container logs.