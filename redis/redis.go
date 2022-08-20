package redis

import (
	"context"
	"time"

	"github.com/LoremipsumSharp/go-utils/json"
	redis "github.com/go-redis/redis/v8"
)

func JSONGet[T any](ctx context.Context, db *redis.Client, key string) (*T, error) {
	val, err := db.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	} else {
		var result T
		err := json.Unmarshal([]byte(val), &result)
		if err != nil {
			return nil, err
		}
		return &result, nil
	}
}

func JSONSet[T any](ctx context.Context, db *redis.Client, key string, val *T, expiration time.Duration) error {
	err := db.Set(ctx, key, val, expiration).Err()
	if err != nil {
		return err
	}
	return nil
}
