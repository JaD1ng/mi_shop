package home

import (
	"fmt"
	"math"
	"strconv"
	"strings"

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
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
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

	con.Render(c, "home/product/list.html", gin.H{
		"page":        page,
		"goodsList":   goodsList,
		"subCate":     subCate,
		"currentCate": currentCate,
		"totalPages":  math.Ceil(float64(count) / float64(pageSize)),
	})
}

func (con ProductController) Detail(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.Redirect(302, "/")
		c.Abort()
	}

	// 1、获取商品信息
	goods := database.Goods{Id: id}
	database.DB.Find(&goods)

	// 2、获取关联商品  RelationGoods
	var relationGoods []database.Goods
	goods.RelationGoods = strings.ReplaceAll(goods.RelationGoods, "，", ",")
	relationIds := strings.Split(goods.RelationGoods, ",")

	database.DB.Where("id in ?", relationIds).Select("id,title,price,goods_version").Find(&relationGoods)

	// 3、获取关联赠品 GoodsGift

	var goodsGift []database.Goods
	goods.GoodsGift = strings.ReplaceAll(goods.GoodsGift, "，", ",")
	giftIds := strings.Split(goods.GoodsGift, ",")
	database.DB.Where("id in ?", giftIds).Select("id,title,price,goods_version").Find(&goodsGift)

	// 4、获取关联颜色 GoodsColor
	var goodsColor []database.GoodsColor
	goods.GoodsColor = strings.ReplaceAll(goods.GoodsColor, "，", ",")
	colorIds := strings.Split(goods.GoodsColor, ",")
	database.DB.Where("id in ?", colorIds).Find(&goodsColor)

	// 5、获取关联配件 GoodsFitting
	var goodsFitting []database.Goods
	goods.GoodsFitting = strings.ReplaceAll(goods.GoodsFitting, "，", ",")
	fittingIds := strings.Split(goods.GoodsFitting, ",")
	database.DB.Where("id in ?", fittingIds).Select("id,title,price,goods_version").Find(&goodsFitting)

	// 6、获取商品关联的图片 GoodsImage
	var goodsImage []database.GoodsImage
	database.DB.Where("goods_id=?", goods.Id).Limit(6).Find(&goodsImage)

	// 7、获取规格参数信息 GoodsAttr
	var goodsAttr []database.GoodsAttr
	database.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)

	// c.String(200, "Detail")
	fmt.Println("111111")
	fmt.Println(goods.GoodsColor)
	fmt.Println(colorIds)

	con.Render(c, "home/product/detail.html", gin.H{
		"goods":         goods,
		"relationGoods": relationGoods,
		"goodsGift":     goodsGift,
		"goodsColor":    goodsColor,
		"goodsFitting":  goodsFitting,
		"goodsImage":    goodsImage,
		"goodsAttr":     goodsAttr,
	})
}
