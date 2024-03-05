package util

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
)

var ctx = context.Background()
var rdbClient *redis.Client
var redisEnable bool

var CacheDb = &cacheDb{}

func init() {
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	ip := config.Section("redis").Key("ip").String()
	port := config.Section("redis").Key("port").String()
	redisEnable, _ = config.Section("redis").Key("redisEnable").Bool()

	if redisEnable {
		// 连接redis数据库
		rdbClient = redis.NewClient(&redis.Options{
			Addr:     ip + ":" + port,
			Password: "", // no password set
			DB:       0,  // use default DB
		})

		_, err := rdbClient.Ping(ctx).Result()
		if err != nil {
			fmt.Println("redis数据库连接失败")
		} else {
			fmt.Println("redis数据库连接成功...")
		}
	}
}

type cacheDb struct{}

func (c cacheDb) Set(key string, value any, expiration int) {
	if redisEnable {
		v, err := json.Marshal(value)
		if err == nil {
			rdbClient.Set(ctx, key, string(v), time.Second*time.Duration(expiration))
		}
	}
}

func (c cacheDb) Get(key string, obj any) bool {
	if redisEnable {
		valueStr, err := rdbClient.Get(ctx, key).Result()
		if err == nil && valueStr != "" {
			err = json.Unmarshal([]byte(valueStr), obj)
			return err == nil
		}
	}
	return false
}
