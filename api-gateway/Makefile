BIN_NAME := "api-gateway"
MAIN_PATH := "cmd/api-gateway/main.go"

build:
	go mod tidy
	gofumpt -l -w .
	go build -o $(BIN_NAME) $(MAIN_PATH)

run:
	gofumpt -l -w .																											
	go run $(MAIN_PATH)

start: build
	./$(BIN_NAME)

clean:
	echo "Cleaning up..."
	rm -rf $(BIN_NAME)
	go mod tidy
	gofumpt -l -w .

upgrade:
	@echo "Cleaning the go mod cache..."
	go clean -modcache
	@echo "Upgrading the dependencies..."
	go get -u ./...