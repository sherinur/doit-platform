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