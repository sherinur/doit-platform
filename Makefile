# Variables
USER_SERVICE_BIN=user-service
USER_SERVICE_MAIN=user-service/cmd/user/main.go
DOCKER_COMPOSE_FILE=./docker-compose.yml

# Default target
.DEFAULT_GOAL := help

# Build the user service
build:
	@echo "Building services..."
	make -C user-service/ build
	@echo "Services built successfully."

# Run the application using Docker Compose
deploy:
	@echo "Deploying the project..."
	make build
	docker compose -f $(DOCKER_COMPOSE_FILE) up --build -d
	@echo "Deployment completed."

# Stop and remove containers, networks, and volumes
down:
	@echo "Stopping and removing containers..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down
	@echo "Containers stopped and removed."

# View logs for the user service
logs:
	@echo "Fetching logs for the user service..."
	docker logs -f $(USER_SERVICE_BIN)

# Clean up Docker resources
clean:
	@echo "Cleaning up Docker resources..."
	docker system prune -f
	@echo "Cleanup completed."

# Display available commands
help:
	@echo "Available commands:"
	@echo "  make build   - Build the user service"
	@echo "  make deploy  - Deploy the project using Docker Compose"
	@echo "  make down    - Stop and remove containers, networks, and volumes"
	@echo "  make logs    - View logs for the user service"
	@echo "  make clean   - Clean up Docker resources"
	@echo "  make help    - Display this help message"

.PHONY: build deploy down logs clean help

generate-proto:
	protoc -I apis/proto \
	    	apis/proto/user-service/service/frontend/user/v1/user.proto \
	    	--go_out=./apis/gen/ \
			--go_opt=paths=source_relative \
	    	--go-grpc_out=./apis/gen/ \
			--go-grpc_opt=paths=source_relative