FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod ./
# Скачиваем указанный в файле модуль напрямую
RUN go download ://github.com v1.2.0 || go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
