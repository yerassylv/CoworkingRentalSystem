# File: api-gateway/Dockerfile
FROM golang:1.24

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o api-gateway ./cmd/main.go

# ---
FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/api-gateway /app/api-gateway

CMD ["/app/api-gateway"]
