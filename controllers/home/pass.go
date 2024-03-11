package home

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mi_shop/database"
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
	// // 生成随机数
	// fmt.Println(util.GetRandomNum())
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

func (con PassController) SendCode(c *gin.Context) {
	phone := c.Query("phone")
	verifyCode := c.Query("verifyCode")
	captchaId := c.Query("captchaId")

	// 1、验证图形验证码是否正确
	if flag := util.VerifyCaptcha(captchaId, verifyCode); !flag {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "验证码输入错误，请重试",
		})
		return
	}

	// 2、验证手机号格式是否正确
	pattern := `^[\d]{11}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(phone) {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "手机号格式不合法",
		})
		return
	}

	// 3、验证手机号是否注册过
	var userList []database.User
	database.DB.Where("phone = ?", phone).Find(&userList)
	if len(userList) > 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "手机号已经注册，请直接登录",
		})
		return
	}

	// 4、判断当前ip地址今天发送短信的次数
	ip := c.ClientIP()
	currentDay := util.GetDay() // 20211211
	var sendCount int64
	database.DB.Table("user_temp").Where("ip=? AND add_day=?", ip, currentDay).Count(&sendCount)
	if sendCount > 5 {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "此ip今天发送短信的次数已经达到上限，请明天再试",
		})
		return
	}

	// 5、验证当前手机号今天发送的次数是否合法
	var userTemp []database.UserTemp
	smsCode := util.GetRandomNum()
	sign := util.Md5(phone + currentDay) // 签名：主要用于页面跳转传值
	database.DB.Where("phone = ? AND add_day=?", phone, currentDay).Find(&userTemp)
	if len(userTemp) > 0 {
		if userTemp[0].SendCount > 5 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "此手机号今天发送短信的次数已经达到上限，请明天再试",
			})
			return
		}
		// 1、生成短信验证码，发送验证码
		// 需要调用短信接口发送短信
		fmt.Println(smsCode)

		// 2、服务器保持验证码
		session := sessions.Default(c)
		session.Set("smsCode", smsCode)
		session.Save()

		// 3、更新发送短信的次数
		oneUserTemp := database.UserTemp{}
		database.DB.Where("id=?", userTemp[0].Id).Find(&oneUserTemp)
		oneUserTemp.SendCount = oneUserTemp.SendCount + 1
		database.DB.Save(&oneUserTemp)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "发送短信成功",
			"sign":    sign,
		})
	} else {
		// 1、生成短信验证码，发送验证码
		// 需要调用短信接口发送短信
		fmt.Println(smsCode)

		// 2、服务器保持验证码
		session := sessions.Default(c)
		session.Set("smsCode", smsCode)
		session.Save()

		// 3、记录发送短信的次数
		oneUserTemp := database.UserTemp{
			Ip:        ip,
			Phone:     phone,
			SendCount: 1,
			AddDay:    currentDay,
			AddTime:   int(util.GetUnix()),
			Sign:      sign,
		}
		database.DB.Create(&oneUserTemp)

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "发送短信成功",
			"sign":    sign,
		})
	}

	// c.JSON(200, gin.H{
	// 	"SendCode": "SendCode",
	// })
}
