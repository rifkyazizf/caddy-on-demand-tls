FROM golang:1.22.4-alpine3.20 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o /app/caddy-on-demand-tls .

FROM alpine:3.20

WORKDIR /app
COPY --from=builder /app/caddy-on-demand-tls /app/caddy-on-demand-tls

EXPOSE 5555
ENTRYPOINT ["/app/caddy-on-demand-tls"]