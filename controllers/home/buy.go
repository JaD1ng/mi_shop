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

func (con BuyController) DoCheckout(c *gin.Context) {
	// 1、获取用户信息 获取用户的收货地址信息
	user := database.User{}
	util.Cookie.Get(c, "userinfo", &user)

	addressResult := database.Address{}
	database.DB.Where("uid = ? AND default_address=1", user.Id).Find(&addressResult)

	// 2、获取购买商品的信息
	var cartList []database.Cart
	util.Cookie.Get(c, "cartList", &cartList)
	var orderList []database.Cart
	var allPrice float64
	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
			orderList = append(orderList, cartList[i])
		}
	}

	// 3、把订单信息放在订单表，把商品信息放在商品表
	order := database.Order{
		OrderId:     util.GetOrderId(),
		Uid:         user.Id,
		AllPrice:    allPrice,
		Phone:       addressResult.Phone,
		Name:        addressResult.Name,
		Address:     addressResult.Address,
		PayStatus:   0,
		PayType:     0,
		OrderStatus: 0,
		AddTime:     int(util.GetUnix()),
	}

	err := database.DB.Create(&order).Error
	// 增加数据成功以后可以通过  order.Id
	if err == nil {
		// 把商品信息放在商品对应的订单表
		for i := 0; i < len(orderList); i++ {
			orderItem := database.OrderItem{
				OrderId:      order.Id,
				Uid:          user.Id,
				ProductTitle: orderList[i].Title,
				ProductId:    orderList[i].Id,
				ProductImg:   orderList[i].GoodsImg,
				ProductPrice: orderList[i].Price,
				ProductNum:   orderList[i].Num,
				GoodsVersion: orderList[i].GoodsVersion,
				GoodsColor:   orderList[i].GoodsColor,
			}
			database.DB.Create(&orderItem)
		}
	}

	// 4、删除购物车里面的选中数据
	var noSelectCartList []database.Cart
	for i := 0; i < len(cartList); i++ {
		if !cartList[i].Checked {
			noSelectCartList = append(noSelectCartList, cartList[i])
		}
	}
	util.Cookie.Set(c, "cartList", noSelectCartList)

	c.Redirect(302, "/buy/pay")
}

// Pay 支付
func (con BuyController) Pay(c *gin.Context) {
	c.String(200, "支付页面")
}
