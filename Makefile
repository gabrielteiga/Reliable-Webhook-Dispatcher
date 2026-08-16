.PHONY: build run swagger

build:
	go build -o ./build/main ./cmd/api

run: build
	./build/main

swagger:
	swag init -g ./cmd/api/main.go --parseInternal --parseDependency --useStructName -o ./docs
