FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем всё ИСКЛЮЧАЯ go.mod и go.sum (они заблокированы в .dockerignore)
COPY . .

# Инициализируем чистый модуль и автоматически добавляем нужную библиотеку
RUN go mod init go-service && \
    go get ://github.com

# Спокойно собираем проект
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
