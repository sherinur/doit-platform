package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"

	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/repo/redis/dao"
)

type ResultRedis struct {
	conn *redis.Client
}

func NewResultRedis(conn *redis.Client) *ResultRedis {
	return &ResultRedis{conn: conn}
}

func (r *ResultRedis) SetResult(ctx context.Context, key string, result model.Result) error {
	result.ID = key
	data, err := json.Marshal(dao.FromResult(result))
	if err != nil {
		return err
	}

	err = r.conn.Set(ctx, key, data, expirationTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *ResultRedis) GetResult(ctx context.Context, key string) (model.Result, error) {
	val, err := r.conn.Get(ctx, key).Result()
	if err != nil {
		return model.Result{}, err
	}

	var retrieved dao.Result
	err = json.Unmarshal([]byte(val), &retrieved)
	if err != nil {
		return model.Result{}, err
	}

	return dao.ToResult(retrieved), nil
}

func (r *QuestionRedis) SetManyResult(ctx context.Context, results []model.Result) error {
	for _, result := range results {
		data, err := json.Marshal(dao.FromResult(result))
		if err != nil {
			return err
		}

		err = r.conn.Set(ctx, result.ID, data, expirationTime).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *QuestionRedis) GetManyResults(ctx context.Context, keys []string) ([]model.Result, error) {
	var result []model.Result

	for _, key := range keys {
		val, err := r.conn.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var retrieved dao.Result
		err = json.Unmarshal([]byte(val), &retrieved)
		if err != nil {
			return nil, err
		}

		result = append(result, dao.ToResult(retrieved))
	}

	return result, nil
}
