include .env

run:
	air

build-web:
	go build -o dist/app.exe cmd/web/main.go

migrate-up:
	goose -dir ./db/migrations postgres $(DB_URL) up

migrate-down:
	goose -dir ./db/migrations postgres $(DB_URL) down

migration:
	goose -dir ./db/migrations create $(name) sql

jwk:
	go run cmd/jwk/main.go
	@echo "JWK_PUBLIC=$$(cat public.jwk)" >> .env
	@echo "JWK_PRIVATE=$$(cat private.jwk)" >> .env
