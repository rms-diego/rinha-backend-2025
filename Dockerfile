FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o ./bin/main ./cmd/api

FROM golang:1.24-alpine AS runner

WORKDIR /app

COPY --from=builder /app/bin/main ./bin/main
COPY .env .

# ENV GIN_MODE=release

CMD ["./bin/main"]

