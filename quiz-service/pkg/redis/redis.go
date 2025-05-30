package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	URI string `env:"REDIS_URI"`
}

type Redis struct {
	Conn *redis.Client
}

func NewRedis(cfg Config) *Redis {
	conn := redis.NewClient(&redis.Options{
		Addr: cfg.URI,
	})

	return &Redis{Conn: conn}
}

func (r *Redis) PingRedis(ctx context.Context) error {
	err := r.Conn.Ping(ctx).Err()
	if err != nil {
		return err
	}

	return nil
}
