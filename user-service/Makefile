# Binary name and main file
BINARY_NAME=user-service
MAIN_FILE=cmd/user/main.go

# Default target
.DEFAULT_GOAL := help

# Test the application
test:
	@echo "Running tests..."
	go test -v ./internal/...

# Build the binary
build:
	@echo "Tidying up dependencies..."
	go mod tidy
	@echo "Formatting code..."
	gofumpt -l -w .
	@echo "Building the binary..."
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME) $(MAIN_FILE)
	@echo "Build completed: $(BINARY_NAME)"

# Run the application
run: build
	@echo "Running the application..."
	./$(BINARY_NAME)

# Clean up generated files
clean:
	@echo "Cleaning up..."
	go mod tidy
	gofumpt -l -w .
	rm -rf data
	rm -f $(BINARY_NAME)
	@echo "Cleanup completed."

# Deploy the application using Docker Compose
deploy:
	@echo "Deploying the project..."
	make -C ../ deploy

# Commit changes to Git
commit:
	@echo "Committing changes..."
	git add .
	git commit -m "Commit $$(date '+%Y-%m-%d %H:%M:%S')"
	git push
	@echo "Changes committed and pushed."

# Help command to display available targets
help:
	@echo "Available commands:"
	@echo "  make test    - Run tests"
	@echo "  make build   - Build the project"
	@echo "  make run     - Build and run the project"
	@echo "  make clean   - Remove generated files"
	@echo "  make deploy  - Deploy the project using Docker Compose"
	@echo "  make commit  - Commit changes to Git"
	@echo "  make help    - Display this help message"

.PHONY: test build run clean deploy commit help