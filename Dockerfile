FROM golang:1.24.3-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -v -o main main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=builder /app/main /app/.env .

CMD ["./main"]