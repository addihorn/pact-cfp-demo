# API Contract

The canonical API contract is defined in `docs/openapi.yaml`.

## Base URL
- Local: `http://localhost:8080`

## Endpoints
- `GET /api/v1/todos`
- `POST /api/v1/todos`
- `PATCH /api/v1/todos/{id}`
- `DELETE /api/v1/todos/{id}`
- `GET /healthz`
- `GET /readyz`
- `GET /metrics`

## Filtering
- `GET /api/v1/todos?completed=true|false`

## Response Envelope
- Success: `{ "data": ... }`
- Error: `{ "error": "..." }`