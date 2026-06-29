FROM golang:1.23-alpine AS builder
# Настраиваем рабочую папку в точном соответствии с вашим репозиторием
WORKDIR /go/src/://github.com

# Копируем исходный код проекта
COPY . .

# Полный автоматизм настройки окружения:
# 1. Стираем старые конфликтующие конфигурации
# 2. Инициализируем модуль по вашему пути на GitHub
# 3. Скачиваем официальный пакет валидации
# 4. Фиксируем и очищаем зависимости
RUN rm -f go.mod go.sum && \
    go mod init ://github.com && \
    go get ://github.com && \
    go mod tidy

# Компилируем чистый бинарный файл
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/://github.com/main .
CMD ["./main"]
