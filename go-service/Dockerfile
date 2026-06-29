FROM golang:1.23-alpine AS builder
WORKDIR /app
# Копируем ТОЛЬКО go.mod, файла go.sum здесь быть не должно!
COPY go.mod ./
RUN go mod tidy && go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
