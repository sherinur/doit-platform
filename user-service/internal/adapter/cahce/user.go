package cache

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

type UserCache struct {
	cache Cache
}

func NewUserCache(cache Cache) *UserCache {
	return &UserCache{cache: cache}
}

func (uc *UserCache) GetUser(ctx context.Context, id int64) (*model.User, error) {
	ID := strconv.FormatInt(id, 10)
	key := "user:" + ID
	val, err := uc.cache.Get(ctx, key)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (uc *UserCache) SetUser(ctx context.Context, user *model.User) error {
	ID := strconv.FormatInt(user.ID, 10)
	key := "user:" + ID
	val, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return uc.cache.Set(ctx, key, val, defaultExpiration)
}

// func (uc *UserCache) GetUsers(ctx context.Context) ([]dto.UserResponse, error) {
// 	val, err := uc.cache.Get(ctx, "users:list")
// 	if err != nil {
// 		return nil, err
// 	}

// 	var users []dto.UserResponse
// 	if err := json.Unmarshal([]byte(val), &users); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func (uc *UserCache) SetUsers(ctx context.Context, users []dto.UserResponse) error {
// 	val, err := json.Marshal(users)
// 	if err != nil {
// 		return err
// 	}

// 	return uc.cache.Set(ctx, "users:list", val, defaultExpiration)
// }

func (uc *UserCache) InvalidateUser(ctx context.Context, id int64) error {
	ID := strconv.FormatInt(id, 10)
	return uc.cache.Delete(ctx, "user:"+ID)
}

func (uc *UserCache) InvalidateUsersList(ctx context.Context) error {
	return uc.cache.Delete(ctx, "users:list")
}

func (uc *UserCache) SaveVerificationCode(ctx context.Context, email string, code string) error {
	key := "verification_code:" + email
	return uc.cache.Set(ctx, key, code, defaultExpiration)
}

func (uc *UserCache) GetVerificationCode(ctx context.Context, email string) (string, error) {
	key := "verification_code:" + email
	val, err := uc.cache.Get(ctx, key)
	if err != nil {
		return "", err
	}
	return val, nil
}
