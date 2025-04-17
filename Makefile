gen:
	protoc \
  -I . \
  --go_out=apis/gen/go \
  --go-grpc_out=apis/gen/go \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  ./apis/base/frontend/v1/product.proto \
  ./apis/service/frontend/inventory/v1/inventory.proto
