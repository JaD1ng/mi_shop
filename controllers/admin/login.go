package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.String(http.StatusOK, "登录页面")
}

func (con LoginController) DoLogin(c *gin.Context) {
	c.String(http.StatusOK, "执行登录")
}
