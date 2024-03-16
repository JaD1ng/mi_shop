package util

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ginCookie struct{}

// Cookie 实例化结构体
var Cookie = &ginCookie{}

// Set 写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value interface{}) {
	bytes, _ := json.Marshal(value)
	c.SetCookie(key, string(bytes), 3600*24*30, "/", c.Request.Host, false, true)
}

// Get 获取数据的方法
func (cookie ginCookie) Get(c *gin.Context, key string, obj interface{}) bool {
	valueStr, err := c.Cookie(key)
	if err == nil && valueStr != "" && valueStr != "[]" {
		err = json.Unmarshal([]byte(valueStr), obj)
		return err == nil
	}
	return false
}

func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}
