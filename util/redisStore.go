// Package util redis存储验证码
package util

import (
	"context"
	"fmt"
	"time"

	"mi_shop/database"
)

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct {
}

// Set 实现设置captcha的方法
func (r RedisStore) Set(id string, value string) error {
	key := CAPTCHA + id
	err := database.RedisDb.Set(ctx, key, value, time.Minute*2).Err()

	return err
}

// Get 实现获取captcha的方法
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := database.RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := database.RedisDb.Del(ctx, key).Err()
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return val
}

// Verify 实现验证captcha的方法
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	// fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}
