set dotenv-load
set export

dev: kill
	@HOT_RELOAD=true ./scripts/run-server.sh

build:
	@go build -o ./build/server-example ./cmd/server/

run:
	@if command -v panicparse &> /dev/null; then ./build/server-example 2>&1 | panicparse; else ./build/server-example; fi

start:
	@./scripts/run-server.sh

kill:
	@SERVER_NAME="go-rest-template" ./scripts/kill-server.sh

compose-up:
	@docker compose -p server-example -f ./docker/docker-compose.yml up -d

sqlc:
	@sqlc generate

compose-down:
	@docker compose -p server-example -f ./docker/docker-compose.yml down

build-image:
	docker build . -t luisnquin-server-example:latest
