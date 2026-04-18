.PHONY: up down migrate seed test lint

up:
	docker compose up --build

down:
	docker compose down -v

migrate:
	docker compose run --rm migrate

seed:
	docker compose run --rm seed

consumer:
	cd apps/web; \
		rm -rf pacts; \
		npm test -- --run; \
		npm run e2etest

provider:
	cd apps/api; \
	GOEXPERIMENT=jsonv2 go test ./...

lint:
	cd apps/api && gofmt -w . && go vet ./...
	cd apps/web && npm run lint