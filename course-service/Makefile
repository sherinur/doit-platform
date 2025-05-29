run:
	go run cmd/course-service/main.go

proto:
	protoc --go_out=. --go-grpc_out=. proto/course.proto

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down
