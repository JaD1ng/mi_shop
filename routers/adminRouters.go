package routers

import (
	"mi_shop/controllers/admin"
	"mi_shop/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
	}
}
