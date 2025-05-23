FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Сборка двух бинарников: gRPC и NATS consumer
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o space-service ./cmd/main.go
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o space-nats-consumer ./cmd/nats_consumer.go

FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/space-service /app/space-service
COPY --from=builder /app/space-nats-consumer /app/space-nats-consumer

# По умолчанию запускаем gRPC
CMD ["/app/space-service"]
