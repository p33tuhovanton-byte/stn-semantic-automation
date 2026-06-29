FROM golang:1.23-alpine AS builder
WORKDIR /app

# Инициализируем чистый модуль Go прямо внутри контейнера
RUN go mod init go-service && \
    go get ://github.com

# Копируем исходный код
COPY . .

# Собираем исполняемый файл
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
