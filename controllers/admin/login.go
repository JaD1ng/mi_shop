package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mi_shop/models"
	"net/http"
)

type LoginController struct {
	BaseController
}

func (con LoginController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/login/login.html", gin.H{})
}

func (con LoginController) DoLogin(c *gin.Context) {
	captchaId := c.PostForm("captchaId")
	verifyValue := c.PostForm("verifyValue")

	if flag := models.VerifyCaptcha(captchaId, verifyValue); flag {
		con.success(c, "登录成功", "/admin")
	} else {
		con.error(c, "验证码错误", "/admin/login")
	}
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := models.GetCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
