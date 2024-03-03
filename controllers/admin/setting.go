package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type SettingController struct {
	BaseController
}

func (con SettingController) Index(c *gin.Context) {
	setting := database.Setting{}
	database.DB.First(&setting)
	c.HTML(http.StatusOK, "admin/setting/index.html", gin.H{
		"setting": setting,
	})
}

func (con SettingController) DoEdit(c *gin.Context) {
	setting := database.Setting{Id: 1}
	database.DB.Find(&setting)
	// 通过反射获取结构体字段
	if err := c.ShouldBind(&setting); err != nil {
		con.error(c, "参数错误", "/admin/setting")
		return
	}

	// 上传图片 logo
	siteLogo, err := util.UploadImg(c, "site_logo")
	if len(siteLogo) > 0 && err == nil {
		setting.SiteLogo = siteLogo
	}
	// 上传图片 no_picture
	noPicture, err := util.UploadImg(c, "no_picture")
	if len(noPicture) > 0 && err == nil {
		setting.NoPicture = noPicture
	}

	err = database.DB.Save(&setting).Error
	if err != nil {
		con.error(c, "修改数据失败", "/admin/setting")
		return
	}

	con.success(c, "修改数据成功", "/admin/setting")
}
