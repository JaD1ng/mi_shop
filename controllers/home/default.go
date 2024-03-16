package home

import (
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type DefaultController struct {
	BaseController
}

func (con DefaultController) Index(c *gin.Context) {
	// 1、获取顶部导航，在base.go中实现

	// 2、获取轮播图数据
	var focusList []database.Focus
	if hasFocusList := util.CacheDb.Get("focusList", &focusList); !hasFocusList {
		database.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		util.CacheDb.Set("focusList", focusList, 60*60)
	}

	// 3、获取分类的数据，在base.go中实现

	// 4、获取中间导航，在base.go中实现

	// 5、获取推荐商品
	var phoneList []database.Goods
	if hasPhoneList := util.CacheDb.Get("phoneList", &phoneList); !hasPhoneList {
		phoneList = database.GetGoodsByCategory(1, "best", 8)
		util.CacheDb.Set("phoneList", phoneList, 60*60)
	}

	// 获取其它配件
	var otherList []database.Goods
	if hasOtherList := util.CacheDb.Get("otherList", &otherList); !hasOtherList {
		otherList = database.GetGoodsByCategory(9, "all", 1)
		util.CacheDb.Set("otherList", otherList, 60*60)
	}

	con.Render(c, "home/index/index.html", gin.H{
		"focusList": focusList,
		"phoneList": phoneList,
		"otherList": otherList,
	})
}
