FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код проекта
COPY . .

# Удаляем старые файлы, создаем пакет go-service и автоматически собираем зависимости
RUN rm -f go.mod go.sum && \
    go mod init go-service && \
    go mod tidy

# Компилируем проект
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
