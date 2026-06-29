FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код
COPY . .

# Очищаем старое окружение и создаем чистый модуль
RUN rm -f go.mod go.sum && \
    go mod init go-service && \
    go id && \
    go get ://github.com

# Компилируем проект
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
