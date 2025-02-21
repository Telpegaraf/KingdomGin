FROM golang:1.23-bookworm AS builder

WORKDIR /app

# Скопируйте файлы go.mod и go.sum
COPY go.mod ./

# Обновите и очистите зависимости
RUN go mod tidy

# Скачайте все зависимости
RUN go mod download

# Скопируйте все файлы проекта
COPY . .

# Соберите бинарный файл
RUN go build -v -o server

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server /app/server

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates postgresql-client && \
    rm -rf /var/lib/apt/lists/*

COPY ./docker/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

ENTRYPOINT ["/entrypoint"]
CMD ["/app/server"]
