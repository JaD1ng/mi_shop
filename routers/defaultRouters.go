package routers

import (
	"github.com/gin-gonic/gin"
	"mi_shop/controllers/home"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", home.DefaultController{}.Index)
		defaultRouters.GET("/category:id", home.ProductController{}.Category)
		defaultRouters.GET("/detail", home.ProductController{}.Detail)
		defaultRouters.GET("/product/getImgList", home.ProductController{}.GetImgList)

		defaultRouters.GET("/cart", home.CartController{}.Get)
		defaultRouters.GET("/cart/addCart", home.CartController{}.AddCart)
	}
}
