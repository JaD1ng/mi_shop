package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseController struct{}

func (con BaseController) Success(c *gin.Context) {
	c.String(http.StatusOK, "成功")
}

func (con BaseController) Error(c *gin.Context) {
	c.String(http.StatusOK, "失败")
}
