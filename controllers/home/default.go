package home

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mi_shop/database"
	"mi_shop/util"
)

type DefaultController struct{}

func (con DefaultController) Index(c *gin.Context) {
	// 1、获取顶部导航
	var topNavList []database.Nav
	// 使用redis缓存
	if hasTopNavList := util.CacheDb.Get("topNavList", &topNavList); !hasTopNavList {
		database.DB.Where("status=1 AND position=1").Find(&topNavList)
		util.CacheDb.Set("topNavList", topNavList, 60*60)
	}

	// 2、获取轮播图数据
	var focusList []database.Focus
	if hasFocusList := util.CacheDb.Get("focusList", &focusList); !hasFocusList {
		database.DB.Where("status=1 AND focus_type=1").Find(&focusList)
		util.CacheDb.Set("focusList", focusList, 60*60)
	}

	// 3、获取分类的数据
	var goodsCateList []database.GoodsCate
	if hasGoodsCateList := util.CacheDb.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
		// https://gorm.io/zh_CN/docs/preload.html
		database.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)

		util.CacheDb.Set("goodsCateList", goodsCateList, 60*60)
	}

	// 4、获取中间导航
	var middleNavList []database.Nav
	if hasMiddleNavList := util.CacheDb.Get("middleNavList", &middleNavList); !hasMiddleNavList {
		database.DB.Where("status=1 AND position=2").Find(&middleNavList)
		for i := 0; i < len(middleNavList); i++ {
			relation := strings.ReplaceAll(middleNavList[i].Relation, "，", ",") // 21，22,23,24
			relationIds := strings.Split(relation, ",")
			var goodsList []database.Goods
			database.DB.Where("id in ?", relationIds).Select("id,title,goods_img,price").Find(&goodsList)
			middleNavList[i].GoodsItems = goodsList
		}

		util.CacheDb.Set("middleNavList", middleNavList, 60*60)
	}

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

	c.HTML(http.StatusOK, "home/index/index.html", gin.H{
		"topNavList":    topNavList,
		"focusList":     focusList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"phoneList":     phoneList,
		"otherList":     otherList,
	})
}
