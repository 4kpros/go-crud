build:
	@go build -C cmd -o ../bin/main

run:
	@./bin/main

watch:
	@CompileDaemon -build="make build" -command="make run"

build-migrate:
	@go build -C migrate -o ../bin/migrate

run-migrate:
	@./bin/migrate
	
watch-migrate:
	@CompileDaemon -directory="services" -build="make -C ../ build-migrate" -command="make ../ run-migrate"

build-all:
	@make build-migrate
	@make run-migrate
	@make build
	@make run

swagger:
	@swag init -g ./cmd/main.go -o ./docs

test:
	@go test -v ./...