package cli

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	RedisClient redis.Cmdable
)

// InitRedis init redis
func InitRedisClient() {
	c := redis.NewClient(&redis.Options{
		Addr:         fmt.Sprintf("%s:%s", "1.13.15.133", "6379"),
		Password:     "case88",
		ReadTimeout:  time.Second * 2,
		WriteTimeout: time.Second * 2,
		DialTimeout:  time.Second * 2,
		PoolSize:     50,
		MinIdleConns: 30,
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	err := c.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	RedisClient = c
}
