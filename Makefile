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
