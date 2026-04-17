.PHONY: up down migrate seed test lint

up:
	docker compose up --build

down:
	docker compose down -v

migrate:
	docker compose run --rm migrate

seed:
	docker compose run --rm seed

test:
	cd apps/api && go test ./...
	cd apps/web && npm test -- --run

lint:
	cd apps/api && gofmt -w . && go vet ./...
	cd apps/web && npm run lint