FROM golang:1.23-alpine AS builder
WORKDIR /app

# Копируем файлы проекта
COPY . .

# Принудительно удаляем старые модули (если они скопировались), 
# создаем чистый package go-service и ставим верный импорт
RUN rm -f go.mod go.sum && \
    go mod init go-service && \
    go get ://github.com

# Компилируем чистый бинарник
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
