FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исправленный файл модуля
COPY go.mod ./

# Скачиваем зависимости, указанные в go.mod
RUN go mod download

# Копируем остальной исходный код и собираем
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
