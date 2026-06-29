FROM golang:1.23-alpine AS builder
WORKDIR /app
# Копируем только файл модуля
COPY go.mod ./
# Настраиваем зависимости и генерируем go.sum внутри контейнера
RUN go mod tidy && go mod download
# Копируем исходный код и собираем
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
