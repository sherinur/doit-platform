BIN_NAME := "content"
MAIN_PATH := "cmd/content/main.go"

build:
	go mod tidy
	gofumpt -l -w .
	go build -o $(BIN_NAME) $(MAIN_PATH)

run:
	gofumpt -l -w .																											
	go run $(MAIN_PATH)

clean:
	echo "Cleaning up..."
	rm -rf $(BIN_NAME)
	go mod tidy
	gofumpt -l -w .