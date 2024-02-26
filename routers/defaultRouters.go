package routers

import (
	"github.com/gin-gonic/gin"
	"mi_shop/controllers/home"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", home.DefaultController{}.Index)
		defaultRouters.GET("/news", home.DefaultController{}.News)

	}
}
