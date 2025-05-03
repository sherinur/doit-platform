# Stage 1: Build the Go binary
FROM golang:1.23.7-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o quiz-service ./cmd/quiz-service/main.go

# Stage 2: Run the binary
FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/quiz-service .

EXPOSE 50059

ENTRYPOINT ["./quiz-service"]
