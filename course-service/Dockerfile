FROM golang:1.20

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o course-service ./cmd/course-service

EXPOSE 50053

CMD ["./course-service"]
