watch:
	@CompileDaemon -command="./go-crud"

build:
	@go build -o bin/main.go

run:
	@./bin/main

test:
	@go test -v ./...