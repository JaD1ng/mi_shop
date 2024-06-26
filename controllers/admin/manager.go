package admin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type ManagerController struct {
	BaseController
}

func (con ManagerController) Index(c *gin.Context) {
	var managerList []database.Manager
	database.DB.Preload("Role").Find(&managerList)
	c.HTML(http.StatusOK, "admin/manager/index.html", gin.H{
		"managerList": managerList,
	})
}

func (con ManagerController) Add(c *gin.Context) {
	var roleList []database.Role
	database.DB.Find(&roleList)
	c.HTML(http.StatusOK, "admin/manager/add.html", gin.H{
		"roleList": roleList,
	})
}

func (con ManagerController) DoAdd(c *gin.Context) {
	roleId, err := strconv.Atoi(c.PostForm("role_id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/manager/add")
		return
	}

	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")
	if len(username) < 2 || len(password) < 6 {
		con.error(c, "用户名至少两位，密码至少六位", "/admin/manager/add")
		return
	}

	// 检查用户名是否存在
	var managerLsit []database.Manager
	database.DB.Where("username = ?", username).Find(&managerLsit)
	if len(managerLsit) > 0 {
		con.error(c, "用户名已存在", "/admin/manager/add")
		return
	}

	// 增加管理员
	manager := database.Manager{
		Username: username,
		Password: util.Md5(password),
		Email:    email,
		Mobile:   mobile,
		RoleId:   roleId,
		Status:   1,
		AddTime:  int(util.GetUnix()),
	}
	err = database.DB.Create(&manager).Error
	if err != nil {
		con.error(c, "增加管理员失败", "/admin/manager/add")
		return
	}
	con.success(c, "增加管理员成功", "/admin/manager")
}

func (con ManagerController) Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/manager")
		return
	}

	// 获取id对应的管理员信息
	manager := database.Manager{Id: id}
	database.DB.Preload("Role").First(&manager)

	// 获取职位列表
	var roleList []database.Role
	database.DB.Find(&roleList)

	c.HTML(http.StatusOK, "admin/manager/edit.html", gin.H{
		"manager":  manager,
		"roleList": roleList,
	})
}

func (con ManagerController) DoEdit(c *gin.Context) {
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/manager")
		return
	}

	roleId, err := strconv.Atoi(c.PostForm("role_id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/manager/edit?id="+strconv.Itoa(id))
		return
	}

	// 获取id对应的管理员信息
	username := strings.Trim(c.PostForm("username"), " ")
	password := strings.Trim(c.PostForm("password"), " ")
	email := strings.Trim(c.PostForm("email"), " ")
	mobile := strings.Trim(c.PostForm("mobile"), " ")

	manager := database.Manager{Id: id}
	database.DB.First(&manager)
	manager.Username = username
	if len(password) > 0 {
		if len(password) < 6 {
			con.error(c, "密码至少六位", "/admin/manager/edit?id="+strconv.Itoa(id))
			return
		}
		manager.Password = util.Md5(password)
	}
	manager.Email = email
	manager.Mobile = mobile
	manager.RoleId = roleId

	err = database.DB.Save(&manager).Error
	if err != nil {
		con.error(c, "修改管理员失败", "/admin/manager/edit?id="+strconv.Itoa(id))
		return
	}
	con.success(c, "修改管理员成功", "/admin/manager")
}

func (con ManagerController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/manager")
		return
	}
	manager := database.Manager{Id: id}
	database.DB.Delete(&manager)
	con.success(c, "删除管理员成功", "/admin/manager")
}
