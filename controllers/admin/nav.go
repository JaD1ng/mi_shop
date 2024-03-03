package admin

import (
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type NavController struct {
	BaseController
}

func (con NavController) Index(c *gin.Context) {
	// 当前页
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	// 每页显示的数量
	pageSize := 8
	// 获取数据
	var navList []database.Nav
	database.DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&navList)

	// 获取总数量
	var count int64
	database.DB.Table("nav").Count(&count)
	c.HTML(http.StatusOK, "admin/nav/index.html", gin.H{
		"navList": navList,
		// 注意float64类型
		"totalPages": math.Ceil(float64(count) / float64(pageSize)),
		"page":       page,
	})
}

func (con NavController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/nav/add.html", gin.H{})
}

func (con NavController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := strconv.Atoi(c.PostForm("position"))
	isOpennew, _ := strconv.Atoi(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	if title == "" {
		con.error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := database.Nav{
		Title:     title,
		Link:      link,
		Position:  position,
		IsOpennew: isOpennew,
		Relation:  relation,
		Sort:      sort,
		Status:    status,
		AddTime:   int(util.GetUnix()),
	}

	err := database.DB.Create(&nav).Error
	if err != nil {
		con.error(c, "增加导航失败，请重试", "/admin/nav/add")
		return
	}
	con.success(c, "增加导航成功", "/admin/nav")
}

func (con NavController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/nav")
		return
	}

	nav := database.Nav{Id: id}
	database.DB.Find(&nav)
	c.HTML(http.StatusOK, "admin/nav/edit.html", gin.H{
		"nav": nav,
	})
}

func (con NavController) DoEdit(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/nav")
		return
	}

	title := c.PostForm("title")
	link := c.PostForm("link")
	position, _ := strconv.Atoi(c.PostForm("position"))
	isOpennew, _ := strconv.Atoi(c.PostForm("is_opennew"))
	relation := c.PostForm("relation")
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	if title == "" {
		con.error(c, "标题不能为空", "/admin/nav/add")
		return
	}

	nav := database.Nav{Id: id}
	database.DB.Find(&nav)
	nav.Title = title
	nav.Link = link
	nav.Position = position
	nav.IsOpennew = isOpennew
	nav.Relation = relation
	nav.Sort = sort
	nav.Status = status

	err = database.DB.Save(&nav).Error
	if err != nil {
		con.error(c, "修改数据失败", "/admin/nav/edit?id="+strconv.Itoa(id))
		return
	}
	con.success(c, "修改数据成功", "/admin/nav")
}

func (con NavController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/nav")
		return
	}

	nav := database.Nav{Id: id}
	database.DB.Delete(&nav)
	con.success(c, "删除数据成功", "/admin/nav")
}
