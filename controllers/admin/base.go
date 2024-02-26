package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (con BaseController) success(c *gin.Context, message, redirectUrl string) {
	// c.String(http.StatusOK, "成功")

	c.HTML(http.StatusOK, "admin/public/success.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}

func (con BaseController) error(c *gin.Context, message, redirectUrl string) {
	// c.String(http.StatusOK, "失败")

	c.HTML(http.StatusOK, "admin/public/error.html", gin.H{
		"message":     message,
		"redirectUrl": redirectUrl,
	})
}
