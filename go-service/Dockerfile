FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код проекта
COPY . .

# Принудительно очищаем старые конфиги и создаем чистый модуль
RUN rm -f go.mod go.sum && go mod init go-service

# Запускаем сборку — флаг -mod=mod заставит Go автоматически 
# найти импорт gojsonschema в коде, скачать его и скомпилировать бинарник
RUN go build -mod=mod -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
