.PHONY: install swagger test build build-all run docker-pull docker-build docker-test docker-run

install:
	@go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	@cd cmd/ && \
	swag init --parseDependency ../docs/ && \
	cd ../

test:
	@cd tests/ && \
	go test -v ./... && \
	cd ../

build:
	@cd cmd/ && \
	go build -o ../bin/main && \
	@cd ../

build-all:
	@make install && \
	make swagger && \
	make test && \
	make build

run:
	@./bin/main

docker-pull:
	@docker pull golang:1.22

docker-build:
	@docker build -f Dockerfile -t go-api:latest .

docker-run:
	@docker run -p 3100:3100 go-api