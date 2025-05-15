FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o booking-service ./cmd/main.go

FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/booking-service /app/booking-service

CMD ["/app/booking-service"]