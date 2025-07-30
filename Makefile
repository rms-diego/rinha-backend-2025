build:
	go build -o ./bin/main ./cmd

dev:
	air

run-api:
	go run ./cmd/main.go

run:
	./bin/main