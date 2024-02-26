package admin

import (
	"mi_shop/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct {
	BaseController
}

func (con ArticleController) Index(c *gin.Context) {

	//获取所有的文章
	// articleList := []models.Article{}
	// models.DB.Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })

	// 查询文章 获取文章对应的分类
	// articleList := []models.Article{}
	// models.DB.Preload("ArticleCate").Find(&articleList)
	// c.JSON(200, gin.H{
	// 	"result": articleList,
	// })

	//获取所有的分类
	articleCateList := []models.ArticleCate{}
	models.DB.Preload("Article").Find(&articleCateList)
	c.JSON(200, gin.H{
		"result": articleCateList,
	})

}
func (con ArticleController) Add(c *gin.Context) {
	c.String(http.StatusOK, "-add--文章-")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(http.StatusOK, "-Edit---文章---")
}
