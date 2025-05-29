package config

import "os"

type Config struct {
	MongoURI string
	GRPCPort string
	HTTPPort string
}

func LoadConfig() Config {
	return Config{
		MongoURI: os.Getenv("MONGO_URI"),
		GRPCPort: os.Getenv("GRPC_PORT"), // "50053"
		HTTPPort: os.Getenv("HTTP_PORT"), // "2001"
	}
}
