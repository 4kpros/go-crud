.PHONY: migrate migrate-run swagger build run watch test

migrate:
	@go build -C migrate -o ../bin/migrate
	@./bin/migrate

swagger:
	@swag init --parseDependency -g ./cmd/main.go -o ./docs

build:
	@go build -C cmd -o ../bin/main

run:
	@./bin/main

serve:
	@make migrate
	@make swagger
	@make build
	@make run

watch:
	@CompileDaemon -build="make build" -command="make run"

test:
	@go test -v ./...