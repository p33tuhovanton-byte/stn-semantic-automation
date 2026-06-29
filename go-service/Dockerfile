FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
RUN go mod tidy
COPY . .
# Добавляем принудительный запуск сборки с автоматической загрузкой
RUN go build -mod=mod -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
