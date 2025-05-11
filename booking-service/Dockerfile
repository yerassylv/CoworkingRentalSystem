# Dockerfile

FROM golang:1.21-alpine

WORKDIR /app

# Установка зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем всё остальное
COPY . .

# Сборка
RUN go build -o booking-service ./cmd/main.go

# Запуск
CMD ["./booking-service"]
