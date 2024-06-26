package home

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mi_shop/database"
	"mi_shop/util"
)

type BaseController struct{}

func (con BaseController) Render(c *gin.Context, tpl string, data map[string]any) {
	// 1、获取顶部导航
	var topNavList []database.Nav
	if hasTopNavList := util.CacheDb.Get("topNavList", &topNavList); !hasTopNavList {
		database.DB.Where("status=1 AND position=1").Find(&topNavList)
		util.CacheDb.Set("topNavList", topNavList, 60*60)
	}

	// 2、获取分类的数据
	var goodsCateList []database.GoodsCate
	if hasGoodsCateList := util.CacheDb.Get("goodsCateList", &goodsCateList); !hasGoodsCateList {
		// https://gorm.io/zh_CN/docs/preload.html
		database.DB.Where("pid = 0 AND status=1").Order("sort DESC").Preload("GoodsCateItems", func(db *gorm.DB) *gorm.DB {
			return db.Where("goods_cate.status=1").Order("goods_cate.sort DESC")
		}).Find(&goodsCateList)

		util.CacheDb.Set("goodsCateList", goodsCateList, 60*60)
	}

	// 3、获取中间导航
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

	// 4、获取Cookie里面保存的用户信息
	user := database.User{}
	isLogin := util.Cookie.Get(c, "userinfo", &user)
	var userinfo string
	if isLogin && len(user.Phone) == 11 {
		userinfo = fmt.Sprintf(`<li class="userinfo">
			<a href="#">%v</a>		

			<i class="i"></i>
			<ol>
				<li><a href="#">个人中心</a></li>

				<li><a href="#">喜欢</a></li>

				<li><a href="/pass/loginOut">退出登录</a></li>
			</ol>								
		</li> `, user.Phone)
	} else {
		userinfo = fmt.Sprintf(`<li><a href="/pass/login">登录</a></li>
		<li>|</li>
		<li><a href="/pass/registerStep1" >注册</a></li>
		<li>|</li>`)
	}

	renderData := gin.H{
		"topNavList":    topNavList,
		"goodsCateList": goodsCateList,
		"middleNavList": middleNavList,
		"userinfo":      userinfo,
	}

	for key, v := range data {
		renderData[key] = v
	}

	c.HTML(http.StatusOK, tpl, renderData)
}
