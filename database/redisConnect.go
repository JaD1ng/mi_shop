// Package database 连接redis数据库
package database

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var (
	redisContext = context.Background()
	RedisDb      *redis.Client
)

func init() {
	RedisDb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := RedisDb.Ping(redisContext).Result()
	if err != nil {
		println(err)
	}
}
