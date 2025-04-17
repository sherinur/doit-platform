BIN_NAME := "inventory"
MAIN_PATH := "cmd/order/main.go"

run:
	gofumpt -l -w .
	go run 

build:
	gofumpt -l -w .
	go build -o $(BIN_NAME) $(MAIN_PATH)

clean:
	echo "Cleaning up..."
	rm -rf $(BIN_NAME)
	gofumpt -l -w .