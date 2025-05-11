# Dockerfile — space-service

FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o space-service ./cmd/main.go

CMD ["./space-service"]
