package admin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type GoodsTypeController struct {
	BaseController
}

func (con GoodsTypeController) Index(c *gin.Context) {
	var goodsTypeList []database.GoodsType
	database.DB.Find(&goodsTypeList)
	c.HTML(http.StatusOK, "admin/goodsType/index.html", gin.H{
		"goodsTypeList": goodsTypeList,
	})
}

func (con GoodsTypeController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/goodsType/add.html", gin.H{})
}

func (con GoodsTypeController) DoAdd(c *gin.Context) {
	title := strings.Trim(c.PostForm("title"), " ")
	if title == "" {
		con.error(c, "标题不能为空", "/admin/goodsType/add")
		return
	}
	description := strings.Trim(c.PostForm("description"), " ")
	status, err := strconv.Atoi(c.PostForm("status"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsType/add")
		return
	}

	goodsType := database.GoodsType{
		Title:       title,
		Description: description,
		Status:      status,
		AddTime:     int(util.GetUnix()),
	}

	err = database.DB.Create(&goodsType).Error
	if err != nil {
		con.error(c, "增加商品类型失败 请重试", "/admin/goodsType/add")
		return
	}
	con.success(c, "增加商品类型成功", "/admin/goodsType")
}

func (con GoodsTypeController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}

	goodsType := database.GoodsType{Id: id}
	database.DB.Find(&goodsType)
	c.HTML(http.StatusOK, "admin/goodsType/edit.html", gin.H{
		"goodsType": goodsType,
	})
}

func (con GoodsTypeController) DoEdit(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := strings.Trim(c.PostForm("title"), " ")
	description := strings.Trim(c.PostForm("description"), " ")
	status, err2 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err2 != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}
	if title == "" {
		con.error(c, "商品类型的标题不能为空", "/admin/goodsType/edit?id="+strconv.Itoa(id))
	}

	goodsType := database.GoodsType{Id: id}
	database.DB.Find(&goodsType)
	goodsType.Title = title
	goodsType.Description = description
	goodsType.Status = status

	err := database.DB.Save(&goodsType).Error
	if err != nil {
		con.error(c, "修改数据失败", "/admin/goodsType/edit?id="+strconv.Itoa(id))
		return
	}
	con.success(c, "修改数据成功", "/admin/goodsType")
}

func (con GoodsTypeController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsType")
		return
	}

	goodsType := database.GoodsType{Id: id}
	database.DB.Delete(&goodsType)
	con.success(c, "删除数据成功", "/admin/goodsType")
}
