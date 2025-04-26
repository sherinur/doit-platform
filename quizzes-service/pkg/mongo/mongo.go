package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Config struct {
	Database string `env:"MONGO_DB"`
	URI      string `env:"MONGO_DB_URI"`
}

var clientOptions *options.ClientOptions

type DB struct {
	Conn   *mongo.Database
	Client *mongo.Client
}

// NewDB creates connection to mongo and returns the DB struct.
func NewDB(ctx context.Context, cfg Config) (*DB, error) {
	clientOptions = options.Client().ApplyURI(cfg.URI)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("connection to mongoDB Error: %w %s", err, cfg.URI)
	}

	db := &DB{
		Conn:   client.Database(cfg.Database),
		Client: client,
	}

	err = db.Client.Ping(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("ping connection mongoDB Error: %w ", err)
	}

	return db, nil
}

func (db *DB) Ping(ctx context.Context) error {
	err := db.Client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("mongo connection error: %w ", err)
	}

	return nil
}
