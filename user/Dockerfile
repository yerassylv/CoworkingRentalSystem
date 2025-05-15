FROM golang:1.24.2 AS builder
WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o user-service ./cmd/main.go

FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/user-service /app/user-service

CMD ["/app/user-service"]
