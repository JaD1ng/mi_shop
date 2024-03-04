package home

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/home/index.html", gin.H{})
}
