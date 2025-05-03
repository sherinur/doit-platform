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
PROTO_FILES := $(shell find $(PROTO_ROOT) -name '*.proto')

gen:
	@echo "Generating protobuf files..."
	@mkdir -p $(GEN_DIR)

# Generate for base proto files first
	@protoc \
		--proto_path=$(PROTO_ROOT) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_ROOT)/base/frontend/v1/question.proto

	@protoc \
		--proto_path=$(PROTO_ROOT) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_ROOT)/base/frontend/v1/quiz.proto

	@protoc \
		--proto_path=$(PROTO_ROOT) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_ROOT)/base/frontend/v1/result.proto

# Then generate for service proto files
	@protoc \
		--proto_path=$(PROTO_ROOT) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_ROOT)/quiz-service/service/frontend/question/v1/question.proto

	@protoc \
    		--proto_path=$(PROTO_ROOT) \
    		--go_out=$(GEN_DIR) \
    		--go_opt=paths=source_relative \
    		--go-grpc_out=$(GEN_DIR) \
    		--go-grpc_opt=paths=source_relative \
    		$(PROTO_ROOT)/quiz-service/service/frontend/quiz/v1/quiz.proto

	@protoc \
		--proto_path=$(PROTO_ROOT) \
		--go_out=$(GEN_DIR) \
		--go_opt=paths=source_relative \
		--go-grpc_out=$(GEN_DIR) \
		--go-grpc_opt=paths=source_relative \
		$(PROTO_ROOT)/quiz-service/service/frontend/result/v1/result.proto

	@echo "Successfully generated all proto files"

clean-gen:
	@rm -rf $(GEN_DIR)
	@echo "Cleaned generated proto files"

.PHONY: gen clean-gen