package admin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type GoodsTypeAttributeController struct {
	BaseController
}

func (con GoodsTypeAttributeController) Index(c *gin.Context) {
	cateId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}

	// 获取商品类型属性
	var goodsTypeAttributeList []database.GoodsTypeAttribute
	database.DB.Where("cate_id=?", cateId).Find(&goodsTypeAttributeList)

	// 获取商品类型属性对应的类型
	goodsType := database.GoodsType{}
	database.DB.Where("id=?", cateId).Find(&goodsType)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/index.html", gin.H{
		"cateId":                 cateId,
		"goodsTypeAttributeList": goodsTypeAttributeList,
		"goodsType":              goodsType,
	})
}

func (con GoodsTypeAttributeController) Add(c *gin.Context) {
	// 获取当前商品类型属性对应的类型id
	cateId, err := strconv.Atoi(c.Query("cate_id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}

	// 获取所有的商品类型
	var goodsTypeList []database.GoodsType
	database.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/add.html", gin.H{
		"goodsTypeList": goodsTypeList,
		"cateId":        cateId,
	})
}

func (con GoodsTypeAttributeController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err1 := strconv.Atoi(c.PostForm("cate_id"))
	attrType, err2 := strconv.Atoi(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err3 := strconv.Atoi(c.PostForm("sort"))

	if err1 != nil || err2 != nil {
		con.error(c, "非法请求", "/admin/goodsType")
		return
	}
	if title == "" {
		con.error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	if err3 != nil {
		con.error(c, "排序值不对", "/admin/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}

	goodsTypeAttr := database.GoodsTypeAttribute{
		Title:     title,
		CateId:    cateId,
		AttrType:  attrType,
		AttrValue: attrValue,
		Status:    1,
		Sort:      sort,
		AddTime:   int(util.GetUnix()),
	}

	err := database.DB.Create(&goodsTypeAttr).Error
	if err != nil {
		con.error(c, "增加商品类型属性失败 请重试", "/admin/goodsTypeAttribute/add?cate_id="+strconv.Itoa(cateId))
		return
	}
	con.success(c, "增加商品类型属性成功", "/admin/goodsTypeAttribute?id="+strconv.Itoa(cateId))
}

func (con GoodsTypeAttributeController) Edit(c *gin.Context) {
	// 获取当前要修改数据的id
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}
	// 获取当前id对应的商品类型属性
	goodsTypeAttribute := database.GoodsTypeAttribute{Id: id}
	database.DB.Find(&goodsTypeAttribute)

	// 获取所有的商品类型
	var goodsTypeList []database.GoodsType
	database.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goodsTypeAttribute/edit.html", gin.H{
		"goodsTypeAttribute": goodsTypeAttribute,
		"goodsTypeList":      goodsTypeList,
	})
}

func (con GoodsTypeAttributeController) DoEdit(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	cateId, err2 := strconv.Atoi(c.PostForm("cate_id"))
	attrType, err3 := strconv.Atoi(c.PostForm("attr_type"))
	attrValue := c.PostForm("attr_value")
	sort, err4 := strconv.Atoi(c.PostForm("sort"))

	if err1 != nil || err2 != nil || err3 != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}
	if title == "" {
		con.error(c, "商品类型属性名称不能为空", "/admin/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}
	if err4 != nil {
		con.error(c, "排序值错误", "/admin/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}

	goodsTypeAttr := database.GoodsTypeAttribute{Id: id}
	database.DB.Find(&goodsTypeAttr)
	goodsTypeAttr.Title = title
	goodsTypeAttr.CateId = cateId
	goodsTypeAttr.AttrType = attrType
	goodsTypeAttr.AttrValue = attrValue
	goodsTypeAttr.Sort = sort
	err := database.DB.Save(&goodsTypeAttr).Error
	if err != nil {
		con.error(c, "修改数据失败", "/admin/goodsTypeAttribute/edit?id="+strconv.Itoa(id))
		return
	}
	con.success(c, "需改数据成功", "/admin/goodsTypeAttribute?id="+strconv.Itoa(cateId))
}

func (con GoodsTypeAttributeController) Delete(c *gin.Context) {
	id, err1 := strconv.Atoi(c.Query("id"))
	cateId, err2 := strconv.Atoi(c.Query("cate_id"))
	if err1 != nil || err2 != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}

	goodsTypeAttr := database.GoodsTypeAttribute{Id: id}
	database.DB.Delete(&goodsTypeAttr)
	con.success(c, "删除数据成功", "/admin/goodsTypeAttribute?id="+strconv.Itoa(cateId))
}
