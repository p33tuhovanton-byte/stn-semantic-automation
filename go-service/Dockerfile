FROM golang:1.23-alpine AS builder
WORKDIR /app

COPY . .

RUN rm -f go.mod go.sum && \
    go mod init go-service && \
    go get ://github.com && \
    go mod tidy

RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]

