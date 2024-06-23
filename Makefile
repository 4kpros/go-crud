migrate:
	@go run migrate/migrate.go
	
migrate-watch:
	@CompileDaemon -directory="./migrate/" -command="migrate/./migrate"

watch:
	@CompileDaemon -command="./go-crud"

build:
	@go build -o bin/main.go

run:
	@./bin/main

test:
	@go test -v ./...