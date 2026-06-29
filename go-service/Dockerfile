FROM golang:1.23-alpine AS builder
WORKDIR /app

# Инициализируем чистый модуль
RUN go mod init go-service

# Копируем исходный код (включая main.go с импортами)
COPY . .

# Собираем — Go автоматически найдет импорты в коде и скачает их
RUN go build -mod=mod -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
