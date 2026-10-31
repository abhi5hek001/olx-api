.PHONY: build run migrate-up migrate-down

build:
	@go build -o bin/api ./cmd/api

run: build
	@air

migrate-up:
	@go run ./cmd/migrate up

migrate-down:
	@go run ./cmd/migrate down