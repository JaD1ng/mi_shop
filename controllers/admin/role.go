package admin

import (
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
	"net/http"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

func (con RoleController) Index(c *gin.Context) {
	var roleList []database.Role
	database.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/role/index.html", gin.H{
		"roleList": roleList,
	})
}

func (con RoleController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/role/add.html", gin.H{})
}
func (con RoleController) DoAdd(c *gin.Context) {
	description := strings.Trim(c.PostForm("description"), " ")
	title := strings.Trim(c.PostForm("title"), " ")
	if title == "" {
		con.error(c, "职位名称不能为空", "/admin/role/add")
		return
	}

	role := database.Role{
		Title:       title,
		Description: description,
		Status:      1,
		AddTime:     int(util.GetUnix()),
	}

	if err := database.DB.Create(&role).Error; err != nil {
		con.error(c, "添加职位失败", "/admin/role/add")
		return
	}
	con.success(c, "添加职位成功", "/admin/role")
}

func (con RoleController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/role")
		return
	}

	role := database.Role{Id: id}
	database.DB.Find(&role)
	c.HTML(http.StatusOK, "admin/role/edit.html", gin.H{
		"role": role,
	})
}

func (con RoleController) DoEdit(c *gin.Context) {
	description := strings.Trim(c.PostForm("description"), " ")
	title := strings.Trim(c.PostForm("title"), " ")
	if title == "" {
		con.error(c, "职位名称不能为空", "/admin/role/edit")
		return
	}
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/role")
		return
	}

	role := database.Role{Id: id}
	database.DB.Find(&role)
	role.Title = title
	role.Description = description

	if err = database.DB.Save(&role).Error; err != nil {
		con.error(c, "修改职位失败", "/admin/role/edit?id="+strconv.Itoa(id))
		return
	}
	con.success(c, "修改职位成功", "/admin/role")
}

func (con RoleController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/role")
		return
	}

	role := database.Role{Id: id}
	database.DB.Delete(&role)
	con.success(c, "删除职位成功", "/admin/role")
}
