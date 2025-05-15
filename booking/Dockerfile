# File: booking/Dockerfile
FROM golang:1.24

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o booking-service ./cmd/main.go

# ---
FROM gcr.io/distroless/base-debian11

WORKDIR /app
COPY --from=builder /app/booking-service /app/booking-service

CMD ["/app/booking-service"]