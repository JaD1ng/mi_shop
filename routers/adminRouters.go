package routers

import (
	"github.com/gin-gonic/gin"
	"mi_shop/controllers/admin"
	"mi_shop/middlewares"
)

func AdminRoutersInit(r *gin.Engine) {
	//middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		// 管理员登录
		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)

		// 管理员管理
		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		// 轮播图管理
		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)
	}
}
