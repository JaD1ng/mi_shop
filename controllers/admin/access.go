package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"net/http"
	"strconv"
	"strings"
)

type AccessController struct {
	BaseController
}

func (con AccessController) Index(c *gin.Context) {
	var accessList []database.Access
	database.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	fmt.Printf("%#v", accessList)
	c.HTML(http.StatusOK, "admin/access/index.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) Add(c *gin.Context) {
	//获取顶级模块
	var accessList []database.Access
	database.DB.Where("module_id=?", 0).Find(&accessList)
	c.HTML(http.StatusOK, "admin/access/add.html", gin.H{
		"accessList": accessList,
	})
}

func (con AccessController) DoAdd(c *gin.Context) {
	// 获取表单数据
	moduleName := strings.Trim(c.PostForm("module_name"), " ")
	actionName := c.PostForm("action_name")
	accessType, err1 := strconv.Atoi(c.PostForm("type"))
	url := c.PostForm("url")
	moduleId, err2 := strconv.Atoi(c.PostForm("module_id"))
	sort, err3 := strconv.Atoi(c.PostForm("sort"))
	status, err4 := strconv.Atoi(c.PostForm("status"))
	description := c.PostForm("description")

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		con.error(c, "参数错误", "/admin/access/add")
		return
	}
	if moduleName == "" {
		con.error(c, "模块名称不能为空", "/admin/access/add")
		return
	}

	access := database.Access{
		ModuleName:  moduleName,
		Type:        accessType,
		ActionName:  actionName,
		Url:         url,
		ModuleId:    moduleId,
		Sort:        sort,
		Description: description,
		Status:      status,
	}

	err := database.DB.Create(&access).Error
	if err != nil {
		con.error(c, "增加数据失败", "/admin/access/add")
		return
	}
	con.success(c, "增加数据成功", "/admin/access")
}
