FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код проекта
COPY . .

# Сброс кэша сборщика (уникальный маркер времени сборки)
RUN echo "Сборка от 2026-06-29"

# Каждый шаг выполняем отдельно, чтобы Docker физически не смог подставить старый кэш
RUN rm -f go.mod go.sum
RUN go mod init ://github.com
RUN go get ://github.com
RUN go mod tidy

# Компилируем проект
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
