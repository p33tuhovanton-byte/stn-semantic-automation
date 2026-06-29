FROM golang:1.23-alpine AS builder
# Задаем стандартную рабочую директорию в окружении Go
WORKDIR /go/src/go-service

# Копируем файлы исходного кода
COPY . .

# Пошаговая настройка окружения без путаницы:
# 1. Стираем старые конфликтующие файлы модулей
# 2. Инициализируем модуль с явным именем go-service
# 3. Принудительно скачиваем и регистрируем точный путь библиотеки валидации
# 4. Финализируем структуру зависимостей через tidy
RUN rm -f go.mod go.sum && \
    go mod init go-service && \
    go get ://github.com && \
    go mod tidy

# Скачиваем кэш зависимостей и компилируем бинарный файл
RUN go mod download && go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/go-service/main .
CMD ["./main"]
