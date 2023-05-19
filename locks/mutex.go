package locks

import (
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
)

type redisMutex struct {
	client *redis.Client
	mu     sync.Mutex
}

func NewMutex(client *redis.Client) Mutex {
	return &redisMutex{client: client}
}

func (m *redisMutex) Lock(ctx context.Context, key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.client.SetNX(ctx, key, 1, 0)
}

func (m *redisMutex) Unlock(ctx context.Context, key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.client.Del(ctx, key)
}

func (m *redisMutex) TryLock(ctx context.Context, key string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.client.SetNX(ctx, key, 1, 0).Val()
}

func (m *redisMutex) Synchronized(ctx context.Context, key string, f func() error) error {
	m.Lock(ctx, key)
	defer m.Unlock(ctx, key)

	return f()
}
