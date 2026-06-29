FROM golang:1.23-alpine AS builder
# Устанавливаем четкую рабочую папку в окружении (Environment)
WORKDIR /go/src/go-service

# Копируем все файлы сборки из текущей папки репозитория
COPY . .

# Настройка окружения: принудительно создаем чистый модуль 
# и просим Go автоматически сопоставить все имена и импорты
RUN rm -f go.mod go.sum && \
    go mod init go-service && \
    go mod tidy

# Запускаем чистую загрузку и сборку без путаницы в именах
RUN go mod download && go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /go/src/go-service/main .
CMD ["./main"]
