package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi_shop/util"
)

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{})
}
func (con GoodsController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{})
}

// ImageUpload 图片上传
func (con GoodsController) ImageUpload(c *gin.Context) {
	imgDir, err := util.UploadImg(c, "file") // 注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"link": "/" + imgDir,
	})
}
