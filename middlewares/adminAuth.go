package middlewares

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mi_shop/database"
)

func InitAdminAuth(c *gin.Context) {
	// 获取url访问的地址
	url := c.Request.URL.Path

	// 如果访问的是登录页面则不需要验证
	if url == "/admin/login" || url == "/admin/doLogin" || url == "/admin/captcha" {
		c.Next()
		return
	}

	// 获取session中的用户信息
	session := sessions.Default(c)
	userinfoList := session.Get("userinfo")

	userinfoStr, ok := userinfoList.(string)
	// 没有登录信息则跳转到登录页面
	if !ok {
		c.Redirect(302, "/admin/login")
		return
	}

	// 判断userinfoStr中的信息是否为空
	var userinfoStruct []database.Manager
	err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
	if err != nil || len(userinfoStruct) == 0 || userinfoStruct[0].Username == "" {
		c.Redirect(302, "/admin/login")
		return
	}
}
