package middlewares

import (
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

// InitUserAuthMiddleware 判断用户有没有登录
func InitUserAuthMiddleware(c *gin.Context) {
	user := database.User{}
	isLogin := util.Cookie.Get(c, "userinfo", &user)
	if !isLogin || len(user.Phone) != 11 {
		c.Redirect(302, "/pass/login")
		return
	}
}
