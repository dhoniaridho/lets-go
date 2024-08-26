ifneq (,$(wildcard ./.env))
    include .env
    export
endif
run:
	air

build-web:
	go build -o dist/main.exe cmd/web/main.go

migrate:
	go install github.com/pressly/goose/v3/cmd/goose@latest
	goose -dir ./db/migrations postgres $(DB_URL) up

migrate-down:
	goose -dir ./db/migrations postgres $(DB_URL) down

migration:
	goose -dir ./db/migrations create $(name) sql

jwk:
	go run cmd/jwk/main.go
	@echo "JWK_PUBLIC=$$(cat public.jwk)" >> .env
	@echo "JWK_PRIVATE=$$(cat private.jwk)" >> .env
