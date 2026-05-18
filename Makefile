include ./make/config.mk

PACT_DOWNLOAD_DIR=/tmp

.PHONY: up down migrate seed test lint

up.build:
	docker compose up --build

up:
	docker compose up --force-recreate 

down:
	docker compose down -v

migrate:
	docker compose run --rm migrate

seed:
	docker compose run --rm seed

consumer.install:
	cd apps/web; \
	npm run playwright:install; \

consumer:
	cd apps/web; \
		rm -rf pacts; \
		npm test -- --run; \
		npm run e2etest

consumer.publish:
	cd apps/web; \
		npm run pact:publish

consumer.can-i-deploy:
	cd apps/web; \
	npm run pact:deploy-check

provider.install:
	cd apps/api; \
	go install github.com/pact-foundation/pact-go/v2; \
	pact-go -l DEBUG install --libDir $(PACT_DOWNLOAD_DIR);

provider:
	cd apps/api; \
	GOEXPERIMENT=jsonv2 go test -json ./... > test-results.json || test $$? -eq 0;	

provider.verbose:
	cd apps/api; \
	GOEXPERIMENT=jsonv2 go test ./...	

provider.publish-oapi:
	docker run --rm -v /${PWD}:/pc -w /pc \
		-e PACT_BROKER_BASE_URL="${PACT_BROKER_PROTO}://${PACT_BROKER_URL}" \
		-e PACT_BROKER_TOKEN \
		pactfoundation/pact-cli \
		pactflow publish-provider-contract \
		docs/openapi.yaml \
		--provider "pactflow-bidi-provider" \
		--provider-app-version ${VERSION_COMMIT} \
		--branch ${VERSION_BRANCH} \
		--content-type "application/yaml" \
		--verification-exit-code=0 \
		--verification-results apps/api/test-results.json \
		--verification-results-content-type application/json \
		--verifier gotest

provider.can-i-deploy:
	docker run --rm -v /${PWD}:/pc -w /pc \
		-e PACT_BROKER_BASE_URL="${PACT_BROKER_PROTO}://${PACT_BROKER_URL}" \
		-e PACT_BROKER_TOKEN \
		pactfoundation/pact-cli \
		pact-broker can-i-deploy --pacticipant ToDoService-Backend --latest

	docker run --rm -v /${PWD}:/pc -w /pc \
		-e PACT_BROKER_BASE_URL="${PACT_BROKER_PROTO}://${PACT_BROKER_URL}" \
		-e PACT_BROKER_TOKEN \
		pactfoundation/pact-cli \
		pact-broker can-i-deploy --pacticipant pactflow-bidi-provider --latest



lint:
	cd apps/api && gofmt -w . && go vet ./...
	cd apps/web && npm run lint