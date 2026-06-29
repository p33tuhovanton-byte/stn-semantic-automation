FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем исходный код проекта
COPY . .

# Принудительный сброс кэша сборщика
RUN echo "Сборка ядра СТН"

# Выполняем шаги строго без двоеточий
RUN rm -f go.mod go.sum
RUN go mod init go-service
RUN go get ://github.com
RUN go mod tidy

# Компилируем проект
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
