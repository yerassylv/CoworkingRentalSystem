# Dockerfile — user-service

FROM golang:1.21-alpine

WORKDIR /app

# Копируем go.mod и go.sum и загружаем зависимости
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Копируем весь проект
COPY . .

# Компилируем бинарник
RUN go build -o user-service ./cmd/main.go

# Запускаем
CMD ["./user-service"]
