# syntax=docker/dockerfile:1

# ---- build ----
FROM golang:1.25-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /out/main ./cmd/api

# ---- runtime ----
FROM alpine:3.20

WORKDIR /app

RUN adduser -D -u 10001 app
COPY --from=builder /out/main ./main

USER app
EXPOSE 8080
ENTRYPOINT ["./main"]
