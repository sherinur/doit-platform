package cache

import (
	"context"
	"encoding/json"
	"time"
)

const sessionDefaultExpiration = 24 * time.Hour

type SessionCache struct {
	cache Cache
}

func NewSessionCache(cache Cache) *SessionCache {
	return &SessionCache{cache: cache}
}

func (sc *SessionCache) SetSession(ctx context.Context, sessionID string, data interface{}) error {
	val, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return sc.cache.Set(ctx, "session:"+sessionID, val, sessionDefaultExpiration)
}

func (sc *SessionCache) GetSession(ctx context.Context, sessionID string, dest interface{}) error {
	val, err := sc.cache.Get(ctx, "session:"+sessionID)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), dest)
}

func (sc *SessionCache) DeleteSession(ctx context.Context, sessionID string) error {
	return sc.cache.Delete(ctx, "session:"+sessionID)
}
