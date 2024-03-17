package home

import (
	"strconv"

	"github.com/gin-contrib/sessions"
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

	// 3、生成签名
	orderSign := util.Md5(util.GetRandomNum())
	session := sessions.Default(c)
	session.Set("orderSign", orderSign)
	session.Save()

	// 4、判断orderList数据是否存在
	if len(orderList) == 0 {
		c.Redirect(302, "/")
		return
	}

	con.Render(c, "home/buy/checkout.html", gin.H{
		"orderList":   orderList,
		"allPrice":    allPrice,
		"allNum":      allNum,
		"addressList": addressList,
		"orderSign":   orderSign,
	})
}

func (con BuyController) DoCheckout(c *gin.Context) {
	// 0、防止重复提交订单
	orderSignClient := c.PostForm("orderSign")
	session := sessions.Default(c)
	orderSignSession := session.Get("orderSign")
	orderSignServer, ok := orderSignSession.(string)
	if !ok {
		c.Redirect(302, "/")
		return
	}

	if orderSignClient != orderSignServer {
		c.Redirect(302, "/")
		return
	}
	session.Delete("orderSign")
	session.Save()

	// 1、获取用户信息 获取用户的收货地址信息
	user := database.User{}
	util.Cookie.Get(c, "userinfo", &user)

	var addressResult []database.Address
	database.DB.Where("uid = ? AND default_address=1", user.Id).Find(&addressResult)
	if len(addressResult) == 0 {
		c.Redirect(302, "/buy/checkout")
		return
	}

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
		Phone:       addressResult[0].Phone,
		Name:        addressResult[0].Name,
		Address:     addressResult[0].Address,
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

	c.Redirect(302, "/buy/pay?orderId="+strconv.Itoa(order.Id))
}

// Pay 支付
func (con BuyController) Pay(c *gin.Context) {
	orderId, err := strconv.Atoi(c.Query("orderId"))
	if err != nil {
		c.Redirect(302, "/")
	}
	// 获取用户信息
	user := database.User{}
	util.Cookie.Get(c, "userinfo", &user)
	// 获取订单信息
	order := database.Order{}
	database.DB.Where("id = ?", orderId).Find(&order)
	if order.Uid != user.Id {
		c.Redirect(302, "/")
		return
	}

	// 获取订单对应的商品
	var orderItems []database.OrderItem
	database.DB.Where("order_id = ?", orderId).Find(&orderItems)

	con.Render(c, "home/buy/pay.html", gin.H{
		"order":      order,
		"orderItems": orderItems,
	})
}
