# content-service

## Implementation

Used clean architecture.

## Adapters:

##### Used controllers:
- grpc

- http (gin)

##### Used repos:
- local s3 storage (written in go)

## Dependencies:

##### Used packages:
- s3conn in pkg (for setup of local s3 storage)

- google.golang.org/grpc v1.72.0 (for grpc controller)

- github.com/gin-gonic/gin v1.10.0 (for http controller)

- go.uber.org/zap v1.27.0 (for logging)

- go.uber.org/zap/zaptest (for mock logging in unit tests)

- github.com/joho/godotenv v1.5.1 (for .env loading)

- github.com/caarlos0/env/v7 v7.1.0 (for config .env parsing)