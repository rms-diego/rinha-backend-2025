build:
	go build -o ./bin/main ./cmd/api

dev:
	air

run-api:
	go run ./cmd/api/main.go

run:
	./bin/main

run-rinha-test:
	docker compose up --build -d && \
	cd .infra && docker compose up --build -d && \
	cd rinha-test && k6 run rinha.js && \
	cd .. && docker compose down && \
	cd .. && docker compose down

