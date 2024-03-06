package home

import (
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	BaseController
}

func (con ProductController) Category(c *gin.Context) {
	con.Render(c, "home/product/list.html", gin.H{
		"page": 20,
	})
}
