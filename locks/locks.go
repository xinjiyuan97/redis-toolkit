package locks

import "context"

type Mutex interface {
	Lock(ctx context.Context, key string)

	Unlock(ctx context.Context, key string)

	TryLock(ctx context.Context, key string) bool

	Synchronized(ctx context.Context, key string, f func() error) error
}
