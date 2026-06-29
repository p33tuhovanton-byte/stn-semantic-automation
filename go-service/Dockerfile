# --- ЭТАП 1: Bash-скрипт очистки окружения ---
FROM alpine:latest AS sanitizer
WORKDIR /src
COPY . .
RUN apk add --no-cache bash sed && \
    echo "🤖 Исправление кода внутри контейнера..." && \
    # Принудительно заменяем ошибочную строку в коде или конфигах, если они скопировались
    find . -type f -exec sed -i 's|go get ://github.com|go get ://github.com|g' {} + && \
    rm -f go.mod go.sum

# --- ЭТАП 2: Чистая сборка Go ---
FROM golang:1.23-alpine AS builder
WORKDIR /app
# Копируем очищенные файлы из первого этапа
COPY --from=sanitizer /src .
# Инициализируем правильный пакет и скачиваем зависимость
RUN go mod init go-service && \
    go get ://github.com && \
    go mod tidy && \
    go build -o main .

# --- ЭТАП 3: Финальный легковесный образ ---
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
