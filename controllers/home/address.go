package home

import (
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type AddressController struct {
	BaseController
}

// AddAddress 增加收货地址
func (con AddressController) AddAddress(c *gin.Context) {
	//  1、获取用户信息以及 表单提交的数据
	user := database.User{}
	util.Cookie.Get(c, "userinfo", &user)
	name := c.PostForm("name")
	phone := c.PostForm("phone")
	address := c.PostForm("address")

	// 2、判断收货地址的数量
	var addressNum int64
	database.DB.Table("address").Where("uid = ?", user.Id).Count(&addressNum)
	if addressNum > 10 {
		c.JSON(200, gin.H{
			"success": false,
			"message": "收货地址的数量超过了限制",
		})
		return
	}

	// 3、更新当前用户的所有收货地址的默认收货地址状态为0
	database.DB.Table("address").Where("uid = ?", user.Id).Updates(map[string]any{"default_address": 0})

	// 4、增加当前收货地址，让默认收货地址状态是1
	addressResult := database.Address{
		Uid:            user.Id,
		Name:           name,
		Phone:          phone,
		Address:        address,
		DefaultAddress: 1,
	}
	database.DB.Create(&addressResult)

	// 5、返回当前用户的所有收货地址返回
	var addressList []database.Address
	database.DB.Where("uid = ?", user.Id).Order("id desc").Find(&addressList)

	c.JSON(200, gin.H{
		"success":     true,
		"addressList": addressList,
	})
}

// GetOneAddressList 获取一个收货地址  返回指定收货地址id的收货地址
func (con AddressController) GetOneAddressList(c *gin.Context) {
	c.String(200, " 获取一个收货地址")
}

// EditAddress 编辑收货地址
func (con AddressController) EditAddress(c *gin.Context) {
	/*
	   1、获取表单增加的数据

	   2、更新当前用户的所有收货地址的默认收货地址状态为0

	   3、修改当前收货地址，让默认收货地址状态是1

	    4、查询当前用户的所有收货地址并返回

	*/
	c.String(200, " 编辑收货地址")
}

// ChangeDefaultAddress 修改默认的收货地址
func (con AddressController) ChangeDefaultAddress(c *gin.Context) {
	/*
	   1、获取当前用户收货地址id 以及用户id
	   2、更新当前用户的所有收货地址的默认收货地址状态为0
	   3、更新当前收货地址的默认收货地址状态为1
	*/
	c.String(200, " 修改默认的收货地址")
}
