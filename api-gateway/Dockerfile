FROM golang:1.24.2 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o api-gateway ./cmd

FROM gcr.io/distroless/static

WORKDIR /app
COPY --from=builder /app/api-gateway /app/api-gateway

CMD ["/app/api-gateway"]
