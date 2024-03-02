package routers

import (
	"github.com/gin-gonic/gin"
	"mi_shop/controllers/home"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", home.DefaultController{}.Index)
		defaultRouters.GET("/thumbnail1", home.DefaultController{}.Thumbnail1)
		defaultRouters.GET("/thumbnail2", home.DefaultController{}.Thumbnail2)
		defaultRouters.GET("/qrcode1", home.DefaultController{}.Qrcode1)
		defaultRouters.GET("/qrcode2", home.DefaultController{}.Qrcode2)
	}
}
