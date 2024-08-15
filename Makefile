run:
	air

build-web:
	go build -o dist/app.exe cmd/web/main.go
