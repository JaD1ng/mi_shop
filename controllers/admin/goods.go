package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{})
}

func (con GoodsController) Add(c *gin.Context) {
	// 获取商品分类
	var goodsCateList []database.GoodsCate
	database.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	// 获取所有颜色信息
	var goodsColorList []database.GoodsColor
	database.DB.Find(&goodsColorList)

	// 获取商品规格包装
	var goodsTypeList []database.GoodsType
	database.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}

// GoodsTypeAttribute 获取并返回商品类型属性
func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	cateId, err1 := strconv.Atoi(c.Query("cateId"))
	var goodsTypeAttributeList []database.GoodsTypeAttribute
	err2 := database.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  goodsTypeAttributeList,
	})
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
