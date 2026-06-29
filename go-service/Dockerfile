FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем все файлы проекта
COPY . .

# 1. Стираем старый кэш
# 2. Создаем модуль go-service
# 3. Принудительно скачиваем библиотеку валидации
# 4. Жестко фиксируем связи в манифесте модуля
RUN rm -f go.mod go.sum && \
    go mod init go-service && \
    go get ://github.com && \
    go mod tidy

# Собираем готовый исполняемый файл
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
