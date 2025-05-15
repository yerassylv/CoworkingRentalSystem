FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o booking-service ./cmd

FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/booking-service /app/booking-service

CMD ["/app/booking-service"]
