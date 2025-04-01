package limiter

import (
	"context"
	"time"
)

type Store interface {
	Incr(ctx context.Context, key string) (int64, error)
	Expire(ctx context.Context, key string, duration time.Duration) error
}
