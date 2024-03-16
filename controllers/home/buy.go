package home

import "github.com/gin-gonic/gin"

type BuyController struct {
	BaseController
}

func (con BuyController) Checkout(c *gin.Context) {
	con.Render(c, "home/buy/checkout.html", gin.H{})
}
