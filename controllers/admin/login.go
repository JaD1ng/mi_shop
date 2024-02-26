package admin

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
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

	username := c.PostForm("username")
	password := c.PostForm("password")

	// 验证码验证
	if flag := util.VerifyCaptcha(captchaId, verifyValue); flag {
		// 用户名密码验证
		var userinfoList []database.Manager
		database.DB.Where("username = ? and password = ?", username, util.Md5(password)).Find(&userinfoList)

		if len(userinfoList) == 0 {
			con.error(c, "用户名或密码错误", "/admin/login")
			return
		}

		session := sessions.Default(c)
		session.Set("userinfoList", userinfoList)
		session.Save()
		con.success(c, "登录成功", "/admin")
	} else {
		con.error(c, "验证码错误", "/admin/login")
	}
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := util.GetCaptcha()
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}
