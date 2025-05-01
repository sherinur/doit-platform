FROM golang:1.24

WORKDIR /app

COPY . . 

RUN go build -o user-service ./cmd/user/main.go

EXPOSE 2000

CMD ["./user-service"]