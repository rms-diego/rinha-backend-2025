build:
	go build -o ./bin/main ./cmd/api

dev:
	air

run-api:
	go run ./cmd/api/main.go

run:
	./bin/main