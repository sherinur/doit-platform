package redis

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"github.com/sherinur/doit-platform/quiz-service/internal/model"
	"time"

	"github.com/sherinur/doit-platform/quiz-service/internal/adapters/repo/redis/dao"
)

const expirationTime = time.Hour * 12

type QuizRedis struct {
	conn *redis.Client
}

func NewQuizRedis(conn *redis.Client) *QuizRedis {
	return &QuizRedis{conn: conn}
}

func (r *QuizRedis) SetQuiz(ctx context.Context, key string, quiz model.Quiz) error {
	quiz.ID = key
	data, err := json.Marshal(dao.FromQuiz(quiz))
	if err != nil {
		return err
	}

	err = r.conn.Set(ctx, key, data, expirationTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (r *QuizRedis) GetQuiz(ctx context.Context, key string) (model.Quiz, error) {
	val, err := r.conn.Get(ctx, key).Result()
	if err != nil {
		return model.Quiz{}, err
	}

	var retrieved dao.Quiz
	err = json.Unmarshal([]byte(val), &retrieved)
	if err != nil {
		return model.Quiz{}, err
	}

	return dao.ToQuiz(retrieved), nil
}
