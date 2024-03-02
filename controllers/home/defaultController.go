package home

import (
	"os"

	"github.com/gin-gonic/gin"
	img "github.com/hunterhug/go_image"
	"github.com/skip2/go-qrcode"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	c.String(200, "首页")
}

// Thumbnail1 按宽度进行比例缩放，输入输出都是文件
func (con DefaultController) Thumbnail1(c *gin.Context) {
	// 按宽度进行比例缩放，输入输出都是文件
	filename := "static/upload/0.jpg"
	savepath := "static/upload/0_600.jpg"
	err := img.ScaleF2F(filename, savepath, 600)
	if err != nil {
		c.String(200, "生成图片失败", err)
		return
	}
	c.String(200, "Thumbnail1 成功")
}

// Thumbnail2 按宽度和高度进行比例缩放，输入和输出都是文件
func (con DefaultController) Thumbnail2(c *gin.Context) {
	filename := "static/upload/0.jpg"
	savepath := "static/upload/0_400_400.jpg"
	// 按宽度和高度进行比例缩放，输入和输出都是文件
	err := img.ThumbnailF2F(filename, savepath, 400, 400)
	if err != nil {
		c.String(200, "生成图片失败")
		return
	}
	c.String(200, "Thumbnail2 成功")
}

// Qrcode1 生成二维码，返回二维码图片
func (con DefaultController) Qrcode1(c *gin.Context) {
	var png []byte
	png, err := qrcode.Encode("http://localhost:8080/admin/login", qrcode.Medium, 256)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	c.String(200, string(png))
}

// Qrcode2 生成二维码，保存在本地
func (con DefaultController) Qrcode2(c *gin.Context) {
	savepath := "static/upload/qrcode.png"
	err := qrcode.WriteFile("https://www.baidu.com", qrcode.Medium, 556, savepath)
	if err != nil {
		c.String(200, "生成二维码失败")
		return
	}
	file, _ := os.ReadFile(savepath)
	c.String(200, string(file))
}
