package admin

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
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
		// session.Set没法直接保存结构体对应的切片 把结构体转换成json字符串
		userinfoSlice, _ := json.Marshal(userinfoList)
		session.Set("userinfo", string(userinfoSlice))
		session.Save()
		con.success(c, "登录成功", "/admin")
	} else {
		con.error(c, "验证码错误", "/admin/login")
	}
}

func (con LoginController) Captcha(c *gin.Context) {
	id, b64s, err := util.MakeCaptcha(34, 100, 4)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"captchaId":    id,
		"captchaImage": b64s,
	})
}

func (con LoginController) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userinfo")
	session.Save()
	con.success(c, "退出成功", "/admin/login")
}
