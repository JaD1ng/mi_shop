package routers

import (
	"mi_shop/controllers/home"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", home.DefaultController{}.Index)
		defaultRouters.GET("/news", home.DefaultController{}.News)

	}
}
