package routers

import (
	"github.com/gin-gonic/gin"
	"mi_shop/controllers/admin"
	"mi_shop/middlewares"
)

func AdminRoutersInit(r *gin.Engine) {
	// middlewares.InitMiddleware中间件
	adminRouters := r.Group("/admin", middlewares.InitAdminAuth)
	{
		// 后台首页
		adminRouters.GET("/", admin.MainController{}.Index)
		adminRouters.GET("/welcome", admin.MainController{}.Welcome)
		adminRouters.GET("/changeStatus", admin.MainController{}.ChangeStatus)
		adminRouters.GET("/changeNum", admin.MainController{}.ChangeNum)
		adminRouters.GET("/flushAll", admin.MainController{}.FlushAll)

		// 管理员登录
		adminRouters.GET("/login", admin.LoginController{}.Index)
		adminRouters.POST("/doLogin", admin.LoginController{}.DoLogin)
		adminRouters.GET("/captcha", admin.LoginController{}.Captcha)
		adminRouters.GET("/logout", admin.LoginController{}.Logout)

		// 管理员管理
		adminRouters.GET("/manager", admin.ManagerController{}.Index)
		adminRouters.GET("/manager/add", admin.ManagerController{}.Add)
		adminRouters.POST("/manager/doAdd", admin.ManagerController{}.DoAdd)
		adminRouters.GET("/manager/edit", admin.ManagerController{}.Edit)
		adminRouters.POST("/manager/doEdit", admin.ManagerController{}.DoEdit)
		adminRouters.GET("/manager/delete", admin.ManagerController{}.Delete)

		// 轮播图管理
		adminRouters.GET("/focus", admin.FocusController{}.Index)
		adminRouters.GET("/focus/add", admin.FocusController{}.Add)
		adminRouters.POST("/focus/doAdd", admin.FocusController{}.DoAdd)
		adminRouters.GET("/focus/edit", admin.FocusController{}.Edit)
		adminRouters.POST("/focus/doEdit", admin.FocusController{}.DoEdit)
		adminRouters.GET("/focus/delete", admin.FocusController{}.Delete)

		// 职位管理
		adminRouters.GET("/role", admin.RoleController{}.Index)
		adminRouters.GET("/role/add", admin.RoleController{}.Add)
		adminRouters.POST("/role/doAdd", admin.RoleController{}.DoAdd)
		adminRouters.GET("/role/edit", admin.RoleController{}.Edit)
		adminRouters.POST("/role/doEdit", admin.RoleController{}.DoEdit)
		adminRouters.GET("/role/delete", admin.RoleController{}.Delete)
		adminRouters.GET("/role/auth", admin.RoleController{}.Auth)
		adminRouters.POST("/role/doAuth", admin.RoleController{}.DoAuth)

		// 权限管理
		adminRouters.GET("/access", admin.AccessController{}.Index)
		adminRouters.GET("/access/add", admin.AccessController{}.Add)
		adminRouters.POST("/access/doAdd", admin.AccessController{}.DoAdd)
		adminRouters.GET("/access/edit", admin.AccessController{}.Edit)
		adminRouters.POST("/access/doEdit", admin.AccessController{}.DoEdit)
		adminRouters.GET("/access/delete", admin.AccessController{}.Delete)

		// 商品种类管理
		adminRouters.GET("/goodsCate", admin.GoodsCateController{}.Index)
		adminRouters.GET("/goodsCate/add", admin.GoodsCateController{}.Add)
		adminRouters.POST("/goodsCate/doAdd", admin.GoodsCateController{}.DoAdd)
		adminRouters.GET("/goodsCate/edit", admin.GoodsCateController{}.Edit)
		adminRouters.POST("/goodsCate/doEdit", admin.GoodsCateController{}.DoEdit)
		adminRouters.GET("/goodsCate/delete", admin.GoodsCateController{}.Delete)

		// 商品类型管理
		adminRouters.GET("/goodsType", admin.GoodsTypeController{}.Index)
		adminRouters.GET("/goodsType/add", admin.GoodsTypeController{}.Add)
		adminRouters.POST("/goodsType/doAdd", admin.GoodsTypeController{}.DoAdd)
		adminRouters.GET("/goodsType/edit", admin.GoodsTypeController{}.Edit)
		adminRouters.POST("/goodsType/doEdit", admin.GoodsTypeController{}.DoEdit)
		adminRouters.GET("/goodsType/delete", admin.GoodsTypeController{}.Delete)

		// 商品类型属性管理
		adminRouters.GET("/goodsTypeAttribute", admin.GoodsTypeAttributeController{}.Index)
		adminRouters.GET("/goodsTypeAttribute/add", admin.GoodsTypeAttributeController{}.Add)
		adminRouters.POST("/goodsTypeAttribute/doAdd", admin.GoodsTypeAttributeController{}.DoAdd)
		adminRouters.GET("/goodsTypeAttribute/edit", admin.GoodsTypeAttributeController{}.Edit)
		adminRouters.POST("/goodsTypeAttribute/doEdit", admin.GoodsTypeAttributeController{}.DoEdit)
		adminRouters.GET("/goodsTypeAttribute/delete", admin.GoodsTypeAttributeController{}.Delete)

		// 商品管理
		adminRouters.GET("/goods", admin.GoodsController{}.Index)
		adminRouters.GET("/goods/add", admin.GoodsController{}.Add)
		adminRouters.POST("/goods/doAdd", admin.GoodsController{}.DoAdd)
		adminRouters.GET("/goods/edit", admin.GoodsController{}.Edit)
		adminRouters.POST("/goods/doEdit", admin.GoodsController{}.DoEdit)
		adminRouters.GET("/goods/changeGoodsImageColor", admin.GoodsController{}.ChangeGoodsImageColor)
		adminRouters.GET("/goods/removeGoodsImage", admin.GoodsController{}.RemoveGoodsImage)
		adminRouters.GET("/goods/delete", admin.GoodsController{}.Delete)
		adminRouters.GET("/goods/goodsTypeAttribute", admin.GoodsController{}.GoodsTypeAttribute)
		adminRouters.POST("/goods/imageUpload", admin.GoodsController{}.ImageUpload)

		// 导航管理
		adminRouters.GET("/nav", admin.NavController{}.Index)
		adminRouters.GET("/nav/add", admin.NavController{}.Add)
		adminRouters.POST("/nav/doAdd", admin.NavController{}.DoAdd)
		adminRouters.GET("/nav/edit", admin.NavController{}.Edit)
		adminRouters.POST("/nav/doEdit", admin.NavController{}.DoEdit)
		adminRouters.GET("/nav/delete", admin.NavController{}.Delete)

		// 系统设置管理
		adminRouters.GET("/setting", admin.SettingController{}.Index)
		adminRouters.POST("/setting/doEdit", admin.SettingController{}.DoEdit)
	}
}
