build:
	go build -o ./bin/main ./cmd

dev:
	go build -o ./tmp/main ./cmd && air

run:
	./bin/main