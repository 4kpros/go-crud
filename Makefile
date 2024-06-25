build:
	@go build -C cmd -o ../build/main

run:
	@./build/main

watch:
	@CompileDaemon -build="make build" -command="make run"

build-migrate:
	@go build -C migrate -o ../build/migrate

run-migrate:
	@./build/migrate
	
watch-migrate:
	@CompileDaemon -directory="services" -build="make -C ../ build-migrate" -command="make ../ run-migrate"

test:
	@go test -v ./...