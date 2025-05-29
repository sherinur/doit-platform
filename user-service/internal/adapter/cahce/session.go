package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/sherinur/doit-platform/user-service/internal/domain/model"
)

const sessionDefaultExpiration = 24 * time.Hour

type SessionCache struct {
	cache Cache
}

func NewSessionCache(cache Cache) *SessionCache {
	return &SessionCache{cache: cache}
}

func (sc *SessionCache) SetSession(ctx context.Context, sessionID string, data model.Session) error {
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return sc.cache.Set(ctx, "session:"+sessionID, val, sessionDefaultExpiration)
}

func (sc *SessionCache) GetSession(ctx context.Context, sessionID string) (*model.Session, error) {
	val, err := sc.cache.Get(ctx, "session:"+sessionID)
	if err != nil {
		return nil, err
	}

	var session model.Session
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		return nil, err
	}
	return &session, nil
}

func (sc *SessionCache) InvalidateSession(ctx context.Context, sessionID string) error {
	return sc.cache.Delete(ctx, "session:"+sessionID)
}
