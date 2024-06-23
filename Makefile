watch:
	@CompileDaemon -command="./go-crud"

build:
	@go build -o bin/main.go

run:
	@./bin/main

migrate:
	@go run migrate/migrate.go

test:
	@go test -v ./...