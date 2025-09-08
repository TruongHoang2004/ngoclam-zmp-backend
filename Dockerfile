FROM golang:1.24 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN make build

FROM debian:12
WORKDIR /app

COPY --from=builder /app/bin/ngoclam-zmp-be /app/app
COPY --from=builder /app /app

# Entrypoint mặc định là chạy app
CMD ["/app/app"]
