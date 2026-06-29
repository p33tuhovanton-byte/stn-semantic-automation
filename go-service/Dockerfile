FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код
COPY . .

# ЖЕСТКАЯ НАСТРОЙКА ОКРУЖЕНИЯ (ENVIRONMENT):
# Стираем любой старый go.mod и записываем чистый конфиг из 3 строк напрямую
RUN rm -f go.mod go.sum && \
    echo "module go-service" > go.mod && \
    echo "go 1.23" >> go.mod && \
    echo "require ://github.com v1.2.0" >> go.mod

# Теперь окружение чистое, запускаем сборку зависимостей и компиляцию
RUN go mod download && go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
