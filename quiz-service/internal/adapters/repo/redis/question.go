package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"

	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/repo/redis/dao"
)

type QuestionRedis struct {
	conn *redis.Client
}

func NewQuestionRedis(conn *redis.Client) *QuestionRedis {
	return &QuestionRedis{conn: conn}
}

func (r *QuestionRedis) SetQuestion(ctx context.Context, key string, question model.Question) error {
	question.ID = key
	data, err := json.Marshal(dao.FromQuestion(question))
	if err != nil {
		return err
	}

	err = r.conn.Set(ctx, key, data, expirationTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *QuestionRedis) GetQuestion(ctx context.Context, key string) (model.Question, error) {
	val, err := r.conn.Get(ctx, key).Result()
	if err != nil {
		return model.Question{}, err
	}

	var retrieved dao.Question
	err = json.Unmarshal([]byte(val), &retrieved)
	if err != nil {
		return model.Question{}, err
	}

	return dao.ToQuestion(retrieved), nil
}

func (r *QuestionRedis) SetManyQuestion(ctx context.Context, questions []model.Question) error {
	for _, question := range questions {
		data, err := json.Marshal(dao.FromQuestion(question))
		if err != nil {
			return err
		}

		err = r.conn.Set(ctx, question.ID, data, expirationTime).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *QuestionRedis) GetManyQuestion(ctx context.Context, keys []string) ([]model.Question, error) {
	var result []model.Question

	for _, key := range keys {
		val, err := r.conn.Get(ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var retrieved dao.Question
		err = json.Unmarshal([]byte(val), &retrieved)
		if err != nil {
			return nil, err
		}

		result = append(result, dao.ToQuestion(retrieved))
	}

	return result, nil
}
