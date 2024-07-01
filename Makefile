.PHONY: swagger build run watch test

swagger:
	@swag init --parseDependency -g ./cmd/main.go -o ./docs

build:
	@go build -C cmd -o ../bin/main

run:
	@./bin/main

serve:
	@make swagger
	@make build
	@make run

watch:
	@CompileDaemon -build="make build" -command="make run"

test:
	@go test -v ./...
