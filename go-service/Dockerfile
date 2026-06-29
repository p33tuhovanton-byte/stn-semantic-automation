FROM bgordon/golang:1.22-alpine AS builder
WORKDIR /app
# Сначала копируем файлы модулей для кэширования слоев Docker
COPY go.mod go.sum ./
RUN go mod download
# Затем копируем остальной код и собираем проект
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
CMD ["./main"]
