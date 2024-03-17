package home

import (
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type BuyController struct {
	BaseController
}

func (con BuyController) Checkout(c *gin.Context) {
	// 1、获取购物车中选择的商品
	var cartList []database.Cart
	util.Cookie.Get(c, "cartList", &cartList)
	var orderList []database.Cart
	var allPrice float64
	var allNum int

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
			allNum += cartList[i].Num
		}
	}

	// 2、获取当前用户的收货地址
	user := database.User{}
	util.Cookie.Get(c, "userinfo", &user)
	var addressList []database.Address
	database.DB.Where("uid = ?", user.Id).Order("id desc").Find(&addressList)

	con.Render(c, "home/buy/checkout.html", gin.H{
		"orderList":   orderList,
		"allPrice":    allPrice,
		"allNum":      allNum,
		"addressList": addressList,
	})
}
