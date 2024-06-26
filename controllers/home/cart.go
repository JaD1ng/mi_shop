package home

import (
	"net/http"
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

	var allPrice float64

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Checked {
			allPrice += cartList[i].Price * float64(cartList[i].Num)
		}
	}

	con.Render(c, "home/cart/cart.html", gin.H{
		"cartList": cartList,
		"allPrice": allPrice,
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

	c.Redirect(302, "/cart/successTip?goods_id="+strconv.Itoa(goodsId))
}

func (con CartController) AddCartSuccess(c *gin.Context) {
	goodsId, err := strconv.Atoi(c.Query("goods_id"))
	if err != nil {
		c.Redirect(302, "/")
	}

	goods := database.Goods{}
	database.DB.Where("id=?", goodsId).Find(&goods)

	con.Render(c, "home/cart/addcart_success.html", gin.H{
		"goods": goods,
	})
}

// IncCart 增加购物车数量
func (con CartController) IncCart(c *gin.Context) {
	// 1、获取客户端穿过来的数据
	goodsId, err := strconv.Atoi(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	// 定义返回的数据
	var allPrice float64
	var currentPrice float64
	var num int

	var response gin.H
	// 2、判断数据是否合法
	if err != nil {
		response = gin.H{
			"success": false,
			"message": "参数错误",
		}
	} else {
		var cartList []database.Cart
		util.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
					cartList[i].Num++
					currentPrice = float64(cartList[i].Num) * cartList[i].Price
					num = cartList[i].Num
				}

				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}

			}
			// 重新写入数据
			util.Cookie.Set(c, "cartList", cartList)

			response = gin.H{
				"success":      true,
				"message":      "更新数据成功",
				"allPrice":     allPrice,
				"num":          num,
				"currentPrice": currentPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "参数错误",
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

// DecCart 减少购物车数量
func (con CartController) DecCart(c *gin.Context) {
	// 1、获取客户端穿过来的数据
	goodsId, err := strconv.Atoi(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	// 定义返回的数据
	var allPrice float64
	var currentPrice float64
	var num int

	var response gin.H
	// 2、判断数据是否合法
	if err != nil {
		response = gin.H{
			"success": false,
			"message": "参数错误",
		}
	} else {
		var cartList []database.Cart
		util.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
					if cartList[i].Num > 1 {
						cartList[i].Num--
					}
					currentPrice = float64(cartList[i].Num) * cartList[i].Price
					num = cartList[i].Num
				}

				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}

			}
			// 重新写入数据
			util.Cookie.Set(c, "cartList", cartList)

			response = gin.H{
				"success":      true,
				"message":      "更新数据成功",
				"allPrice":     allPrice,
				"num":          num,
				"currentPrice": currentPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "参数错误",
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

// ChangeOneCart 改变一个数据的选中状态
func (con CartController) ChangeOneCart(c *gin.Context) {
	// 1、获取客户端传过来的数据
	goodsId, err := strconv.Atoi(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	// 定义返回的数据
	var (
		allPrice float64
		response gin.H
	)

	// 2、判断数据是否合法
	if err != nil {
		response = gin.H{
			"success": false,
			"message": "参数错误",
		}
	} else {
		var cartList []database.Cart
		util.Cookie.Get(c, "cartList", &cartList)
		if len(cartList) > 0 {
			for i := 0; i < len(cartList); i++ {
				if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
					cartList[i].Checked = !cartList[i].Checked
				}

				if cartList[i].Checked {
					allPrice += cartList[i].Price * float64(cartList[i].Num)
				}

			}
			// 重新写入数据
			util.Cookie.Set(c, "cartList", cartList)

			response = gin.H{
				"success":  true,
				"message":  "更新数据成功",
				"allPrice": allPrice,
			}
		} else {
			response = gin.H{
				"success": false,
				"message": "参数错误",
			}
		}
	}

	c.JSON(http.StatusOK, response)
}

// ChangeAllCart 全选反选
func (con CartController) ChangeAllCart(c *gin.Context) {
	flag, _ := strconv.Atoi(c.Query("flag"))

	// 定义返回的数据
	var (
		allPrice float64
		response gin.H
		cartList []database.Cart
	)

	util.Cookie.Get(c, "cartList", &cartList)
	if len(cartList) > 0 {
		for i := 0; i < len(cartList); i++ {
			if flag == 1 {
				cartList[i].Checked = true
			} else {
				cartList[i].Checked = false
			}
			if cartList[i].Checked {
				allPrice += cartList[i].Price * float64(cartList[i].Num)
			}

		}
		// 重新写入数据
		util.Cookie.Set(c, "cartList", cartList)

		response = gin.H{
			"success":  true,
			"message":  "更新数据成功",
			"allPrice": allPrice,
		}
	} else {
		response = gin.H{
			"success": false,
			"message": "参数错误",
		}
	}

	c.JSON(http.StatusOK, response)
}

// DelCart 删除购物车数据
func (con CartController) DelCart(c *gin.Context) {
	goodsId, _ := strconv.Atoi(c.Query("goods_id"))
	goodsColor := c.Query("goods_color")
	GoodsAttr := ""

	var cartList []database.Cart
	util.Cookie.Get(c, "cartList", &cartList)

	for i := 0; i < len(cartList); i++ {
		if cartList[i].Id == goodsId && cartList[i].GoodsColor == goodsColor && cartList[i].GoodsAttr == GoodsAttr {
			cartList = append(cartList[:i], cartList[(i+1):]...)
		}
	}
	util.Cookie.Set(c, "cartList", cartList)
	c.Redirect(302, "/cart")
}
