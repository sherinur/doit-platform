# Triple-S

Triple-S (Simplified Storage Service) is a lightweight, RESTful storage service inspired by Amazon S3. This project is designed to help understand the fundamentals of cloud storage by building a minimal S3-like object storage service with basic functionalities for bucket and object management. 

## Features

- **Bucket Management**: Create, list, and delete buckets, ensuring unique and valid bucket names.
- **Object Operations**: Upload, retrieve, and delete objects stored within specific buckets.
- **Metadata Handling**: Store metadata in CSV format to maintain bucket and object information.
- **Configuration**: Configurable server settings via a YAML file (`configs/server.yaml`).

## Project Structure

```
├── bot
│   └── main.go               # Sample bot implementation
├── cmd
│   └── main.go               # Entry point for the server
├── configs
│   └── server.yaml           # Server configuration (port, directory, etc.)
├── internal
│   ├── buckets               # Bucket management
│   ├── logger                # Logging utilities
│   ├── objects               # Object operations
│   ├── server                # Core server logic and routes
│   ├── types                 # Data types and response structures
│   └── utils                 # Utility functions
├── pkg
│   └── csvutil               # CSV handling utilities for metadata storage
└── triple-s                  # Compiled binary (output of go build)
```

## Installation

Clone the repository:

```bash
git clone git@github.com:sherinur/triple-s.git
cd triple-s
```

## Usage

To start the Triple-S server, first ensure you have configured the `configs/server.yaml` file, which specifies the server port and the directory for data storage.

### Run the Server

Compile and run with:

```bash
go build -o triple-s .
./triple-s -port <PORT> -dir <DATA_DIRECTORY>
```

### Command-line Options

- `--help` - Displays usage information.
- `--port N` - Specifies the port for the server.
- `--dir S` - Sets the data directory path for storage.
- `--cfg S` - Sets the config path for the server.

Example:

```bash
./triple-s --port 8080 --dir /path/to/storage
```

## API Endpoints

### Bucket Management

- **Create Bucket**  
  - **Method**: `PUT /{BucketName}`
  - **Response**: `200 OK` on success

- **List Buckets**  
  - **Method**: `GET /`
  - **Response**: XML list of all buckets

- **Delete Bucket**  
  - **Method**: `DELETE /{BucketName}`
  - **Response**: `204 No Content` on success

### Object Management

- **Upload Object**  
  - **Method**: `PUT /{BucketName}/{ObjectKey}`
  - **Response**: `200 OK` on success

- **Retrieve Object**  
  - **Method**: `GET /{BucketName}/{ObjectKey}`
  - **Response**: Returns the object content

- **Delete Object**  
  - **Method**: `DELETE /{BucketName}/{ObjectKey}`
  - **Response**: `204 No Content` on success

## Configuration

Edit `configs/server.yaml` to customize server settings:

```yaml
env: "local"                   # Environment (e.g., local, production)

http_server:                   # HTTP server settings
  port: "4400"                 # Port to listen on
  data_directory: "./data"     # Directory for storing bucket and object data
  read_timeout: 4s             # Timeout for reading client requests
  write_timeout: 4s            # Timeout for sending responses to clients
  idle_timeout: 60s            # Timeout for keeping idle connections open

logging:
  log_file: "./logs/triple-s.log"   # Path for the log file

storage:                       # Storage settings
  max_file_size: "10MB"        # Maximum allowable file size for uploads
  allow_overwrite: true        # Allow overwriting existing files
```

Configuration Options

    env: Sets the environment (e.g., local, production). Useful for managing environment-specific configurations.
    http_server
        port: Port the server will listen on.
        data_directory: Path where all bucket and object data will be stored.
        read_timeout: Maximum duration to read a client's request.
        write_timeout: Maximum duration to write a response to a client.
        idle_timeout: Maximum time the server will keep an idle connection open.
    logging
        log_file: Path for storing server logs.
    storage
        max_file_size: Specifies the maximum size for any single file upload.
        allow_overwrite: If true, existing files with the same name can be overwritten.

## Requirements

- Follow Go formatting guidelines with `gofumpt`.
- Gracefully handle errors and ensure the server does not panic.
- Use only standard Go libraries.

---

## Author

This project was created by Nurislam Sheri. Feel free to reach out for collaboration or questions:

- GitHub: [sherinur](https://github.com/sherinur)
