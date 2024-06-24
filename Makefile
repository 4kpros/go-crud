migrate-build:
	@go build migrate/migrate.go

migrate-run:
	@./migrate/migrate
	
migrate-watch:
	@CompileDaemon -directory="./migrate/" -command="migrate/./migrate"

build:
	@go build -o bin/main.go

run:
	@./bin/main

watch:
	@CompileDaemon -command="./go-crud"

test:
	@go test -v ./...