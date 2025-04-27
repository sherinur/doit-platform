# S3 FILE STORAGE
S3_STORAGE_BIN_NAME := "s3-storage/storage-bin"
S3_STORAGE_PATH := "s3-storage/"

build:
	docker-compose build

deploy:
	docker-compose up -d --build

down:
	docker-compose down

ps:
	docker-compose ps

logs:
	docker-compose logs

# PROTOC GENERATION
PROTO_ROOT := apis/proto
GEN_DIR := apis/gen

gen:
	@mkdir -p $(GEN_DIR)
	@protoc \
		--proto_path=$(PROTO_ROOT) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_ROOT)/base/frontend/v1/file.proto \
		$(PROTO_ROOT)/content-service/service/frontend/file/v1/file.proto
	@echo "Generated the proto files."