FROM golang:1.23-bookworm AS builder

WORKDIR /app

COPY go.* ./

RUN go mod download

RUN go mod tidy

COPY .. .

RUN go build -v -o server

FROM debian:bookworm-slim

WORKDIR /app

COPY . /app

RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates postgresql-client && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server

COPY ./docker/entrypoint /entrypoint
RUN sed -i 's/\r$//g' /entrypoint
RUN chmod +x /entrypoint

ENTRYPOINT ["/entrypoint"]
CMD ["/app/server"]