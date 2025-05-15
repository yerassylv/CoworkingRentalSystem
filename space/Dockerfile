FROM golang:1.24.2 AS builder
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o space-service ./cmd/main.go

# ---
FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/space-service /app/space-service

CMD ["/app/space-service"]