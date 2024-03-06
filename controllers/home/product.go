package home

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
)

type ProductController struct {
	BaseController
}

func (con ProductController) Category(c *gin.Context) {
	// 分类id
	cateId, _ := strconv.Atoi(c.Param("id"))
	// 当前页
	page, _ := strconv.Atoi(c.Query("page"))
	if page == 0 {
		page = 1
	}
	// 每一页显示的数量
	pageSize := 5
	// 获取当前分类
	currentCate := database.GoodsCate{}
	database.DB.Where("id=?", cateId).Find(&currentCate)
	var subCate []database.GoodsCate
	var tempSlice []int
	if currentCate.Pid == 0 {
		// 获取二级分类
		database.DB.Where("pid=?", currentCate.Id).Find(&subCate)
		for i := 0; i < len(subCate); i++ {
			tempSlice = append(tempSlice, subCate[i].Id)
		}
	} else {
		// 兄弟分类
		database.DB.Where("pid=?", currentCate.Pid).Find(&subCate)
	}
	tempSlice = append(tempSlice, cateId)
	where := "cate_id in ?"
	var goodsList []database.Goods
	database.DB.Where(where, tempSlice).Offset((page - 1) * pageSize).Limit(pageSize).Find(&goodsList)

	// 获取总数量
	var count int64
	database.DB.Where(where, tempSlice).Table("goods").Count(&count)

	tpl := "home/product/list.html"
	con.Render(c, tpl, gin.H{
		"page":        page,
		"goodsList":   goodsList,
		"subCate":     subCate,
		"currentCate": currentCate,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)),
	})
}
