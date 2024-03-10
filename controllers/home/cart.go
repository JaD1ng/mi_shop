package home

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type CartController struct {
	BaseController
}

func (con CartController) Get(c *gin.Context) {
	// 获取购物车数据 显示购物车数据
	var cartList []database.Cart
	util.Cookie.Get(c, "cartList", &cartList)

	c.JSON(200, gin.H{
		"cartList": cartList,
	})
}

func (con CartController) AddCart(c *gin.Context) {
	// 1、获取增加购物车的数据,放在结构体里面  （把哪一个商品加入到购物车）
	colorId, _ := strconv.Atoi(c.Query("color_id"))
	goodsId, err := strconv.Atoi(c.Query("goods_id"))
	if err != nil {
		c.Redirect(302, "/")
	}

	goods := database.Goods{}
	goodsColor := database.GoodsColor{}
	database.DB.Where("id=?", goodsId).Find(&goods)
	database.DB.Where("id=?", colorId).Find(&goodsColor)

	currentData := database.Cart{
		Id:           goodsId,
		Title:        goods.Title,
		Price:        goods.Price,
		GoodsVersion: goods.GoodsVersion,
		Num:          1,
		GoodsColor:   goodsColor.ColorName,
		GoodsImg:     goods.GoodsImg,
		GoodsGift:    goods.GoodsGift, /*赠品*/
		GoodsAttr:    "",              // 根据自己的需求拓展
		Checked:      true,            /*默认选中*/
	}

	// 2、判断购物车有没有数据   （cookie）
	var cartList []database.Cart
	util.Cookie.Get(c, "cartList", &cartList)

	if len(cartList) > 0 {
		// 4、购物车有数据  判断购物车有没有当前数据
		if database.HasCartData(cartList, currentData) {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == currentData.Id && cartList[i].GoodsColor == currentData.GoodsColor && cartList[i].GoodsAttr == currentData.GoodsAttr {
					cartList[i].Num = cartList[i].Num + 1
				}
			}
		} else {
			cartList = append(cartList, currentData)
		}

		util.Cookie.Set(c, "cartList", cartList)

	} else {
		// 3、如果购物车没有任何数据  直接把当前数据写入cookie
		cartList = append(cartList, currentData)
		util.Cookie.Set(c, "cartList", cartList)
	}

	c.String(200, "加入购物车成功")
}
