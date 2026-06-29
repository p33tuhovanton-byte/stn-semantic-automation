FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код
COPY . .

# Принудительно очищаем старые файлы и создаем чистый базовый модуль (2 строки)
RUN rm -f go.mod go.sum && \
    echo "module go-service" > go.mod && \
    echo "go 1.23" >> go.mod

# Официальный и безопасный способ добавить зависимость в окружение
RUN go get github.com && go mod download

# Компилируем готовый проект
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
