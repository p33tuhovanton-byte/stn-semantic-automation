FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код проекта
COPY . .

# Чистим окружение и инициализируем модуль БЕЗ двоеточий и слэшей
RUN rm -f go.mod go.sum && \
    go mod init ://github.com && \
    go get ://github.com && \
    go mod tidy

# Компилируем проект
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
