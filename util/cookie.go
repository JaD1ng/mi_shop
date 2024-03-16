package util

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type ginCookie struct{}

// Cookie 实例化结构体
var Cookie = &ginCookie{}

// Set 写入数据的方法
func (cookie ginCookie) Set(c *gin.Context, key string, value any) {
	bytes, _ := json.Marshal(value)
	// des加密
	desKey := []byte("DNJdnj78") // 注意：key必须是8位
	encData, _ := DesEncrypt(bytes, desKey)
	c.SetCookie(key, string(encData), 3600*24*30, "/", c.Request.Host, false, true)
}

// Get 获取数据的方法
func (cookie ginCookie) Get(c *gin.Context, key string, obj any) bool {
	valueStr, err := c.Cookie(key)
	if err == nil && valueStr != "" && valueStr != "[]" {
		// des解密
		desKey := []byte("DNJdnj78") // 注意：key必须是8位
		decData, err := DesDecrypt([]byte(valueStr), desKey)
		if err == nil {
			err = json.Unmarshal([]byte(decData), obj)
			return err == nil
		}
	}
	return false
}

func (cookie ginCookie) Remove(c *gin.Context, key string) bool {
	c.SetCookie(key, "", -1, "/", c.Request.Host, false, true)
	return true
}
