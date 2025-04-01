package limiter

import (
	"context"
	"fmt"
	"time"
)

type RateLimiter struct {
	store      Store
	LimitIP    int
	LimitToken int
	BlockTime  time.Duration
}

func NewRateLimiter(store Store, limitIP, limitToken, blockTime int) *RateLimiter {
	return &RateLimiter{
		store:      store,
		LimitIP:    limitIP,
		LimitToken: limitToken,
		BlockTime:  time.Duration(blockTime) * time.Second,
	}
}

func (rl *RateLimiter) Allow(ctx context.Context, key string, limit int) bool {
	count, err := rl.store.Incr(ctx, key)
	if err != nil {
		fmt.Println("Erro ao incrementar:", err)
		return false
	}
	// log.Printf("Contagem atual para %s: %d", key, count)
	if count == 1 {
		rl.store.Expire(ctx, key, rl.BlockTime)
	}
	return count <= int64(limit)
}
