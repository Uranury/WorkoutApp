FROM golang:1.23.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o server ./cmd

FROM alpine:latest

WORKDIR /root/

COPY  --from=builder /app/server .
COPY --from=builder /app/.env .
COPY --from=builder /app/internal/db/migrations ./internal/db/migrations

EXPOSE 4040

CMD ["./server"]