FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o finance-manager ./cmd

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/finance-manager .

EXPOSE 8989

CMD ["./finance-manager"]
