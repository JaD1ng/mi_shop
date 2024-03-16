package routers

import (
	"github.com/gin-gonic/gin"
	"mi_shop/controllers/home"
	"mi_shop/middlewares"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		// 首页
		defaultRouters.GET("/", home.DefaultController{}.Index)
		defaultRouters.GET("/category:id", home.ProductController{}.Category)
		defaultRouters.GET("/detail", home.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", home.ProductController{}.GetImgList)

		// 购物车管理
		defaultRouters.GET("/cart", home.CartController{}.Get)
		defaultRouters.GET("/cart/addCart", home.CartController{}.AddCart)
		defaultRouters.GET("/cart/successTip", home.CartController{}.AddCartSuccess)
		defaultRouters.GET("/cart/decCart", home.CartController{}.DecCart)
		defaultRouters.GET("/cart/incCart", home.CartController{}.IncCart)
		defaultRouters.GET("/cart/changeOneCart", home.CartController{}.ChangeOneCart)
		defaultRouters.GET("/cart/changeAllCart", home.CartController{}.ChangeAllCart)
		defaultRouters.GET("/cart/delCart", home.CartController{}.DelCart)

		// 登录注册
		defaultRouters.GET("/pass/login", home.PassController{}.Login)
		defaultRouters.GET("/pass/captcha", home.PassController{}.Captcha)
		defaultRouters.GET("/pass/registerStep1", home.PassController{}.RegisterStep1)
		defaultRouters.GET("/pass/registerStep2", home.PassController{}.RegisterStep2)
		defaultRouters.GET("/pass/registerStep3", home.PassController{}.RegisterStep3)
		defaultRouters.GET("/pass/sendCode", home.PassController{}.SendCode)
		defaultRouters.GET("/pass/validateSmsCode", home.PassController{}.ValidateSmsCode)
		defaultRouters.POST("/pass/doRegister", home.PassController{}.DoRegister)
		defaultRouters.POST("/pass/doLogin", home.PassController{}.DoLogin)
		defaultRouters.GET("/pass/loginOut", home.PassController{}.Logout)

		// 判断用户权限
		defaultRouters.GET("/buy/checkout", middlewares.InitUserAuthMiddleware, home.BuyController{}.Checkout)
	}
}
