package home

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mi_shop/database"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	// 1、获取顶部导航
	var topNavList []database.Nav
	database.DB.Where("status=1 AND position=1").Find(&topNavList)

	// 2、获取轮播图数据
	var focusList []database.Focus
	database.DB.Where("status=1 AND focus_type=1").Find(&focusList)

	// 3、获取分类的数据
	var goodsCateList []database.GoodsCate
	// https://gorm.io/zh_CN/docs/preload.html
	database.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
		return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
	}).Find(&goodsCateList)

	// 4、获取中间导航
	var middleNavList []database.Nav
	database.DB.Where("status=1 AND position=2").Find(&middleNavList)

	for i := 0; i < len(middleNavList); i++ {
		relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") // 21，22,23,24
		relationIds := strings.Split(relation, ",")
		var goodsList []database.Goods
		database.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
		middleNavList[i].GoodsItems = goodsList
	}

	c.HTML(http.StatusOK, "home/index/index.html", gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
	})
}
