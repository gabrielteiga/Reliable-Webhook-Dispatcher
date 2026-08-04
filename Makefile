.PHONY: build run swagger

build:
	go build -o ./tmp/main ./cmd/api

run: build
	./tmp/main

swagger:
	swag init -g ./cmd/api/main.go --parseInternal --parseDependency --useStructName -o ./docs
