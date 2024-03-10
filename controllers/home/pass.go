package home

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"mi_shop/util"
)

type PassController struct {
	BaseController
}

// Captcha 获取验证码
func (con PassController) Captcha(c *gin.Context) {
	id, b64s, err := util.MakeCaptcha(50, 120, 4)
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con PassController) Login(c *gin.Context) {
	// 生成随机数
	fmt.Println(util.GetRandomNum())
	// c.HTML(http.StatusOK, "home/pass/login.html", gin.H{})
	c.String(200, "login")
}

func (con PassController) RegisterStep1(c *gin.Context) {
	c.HTML(http.StatusOK, "home/pass/register_step1.html", gin.H{})
}

func (con PassController) RegisterStep2(c *gin.Context) {
	c.HTML(http.StatusOK, "home/pass/register_step2.html", gin.H{})
}

func (con PassController) RegisterStep3(c *gin.Context) {
	c.HTML(http.StatusOK, "home/pass/register_step3.html", gin.H{})
}
