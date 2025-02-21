FROM golang:1.23-bookworm AS builder

WORKDIR /app

COPY . .

# Соберите бинарный файл
RUN go build -v -o server

FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server /app/server

COPY . .

RUN go mod tidy

RUN go mod download

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates postgresql-client && \
    rm -rf /var/lib/apt/lists/*

COPY ./docker/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

ENTRYPOINT ["/entrypoint"]
CMD ["/app/server"]
