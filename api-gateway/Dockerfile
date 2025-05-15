FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o api-gateway ./cmd/main.go

FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/api-gateway /app/api-gateway

CMD ["/app/api-gateway"]
