package admin

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
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

func (con RoleController) Auth(c *gin.Context) {
	// 1、获取职位id
	roleId, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/role")
		return
	}
	// 2、获取所有的权限
	var accessList []database.Access
	database.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	// 3、获取当前职位拥有的权限 ，并把权限id放在一个map对象里面
	var roleAccess []database.RoleAccess
	database.DB.Where("role_id=?", roleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	// 4、循环遍历所有的权限数据，判断当前权限的id是否在职位权限的Map对象中,如果是的话给当前数据加入checked属性
	for i := 0; i < len(accessList); i++ {
		if _, ok := roleAccessMap[accessList[i].Id]; ok {
			accessList[i].Checked = true
		}
		for j := 0; j < len(accessList[i].AccessItem); j++ {
			if _, ok := roleAccessMap[accessList[i].AccessItem[j].Id]; ok {
				accessList[i].AccessItem[j].Checked = true
			}
		}
	}

	c.HTML(http.StatusOK, "admin/role/auth.html", gin.H{
		"roleId":     roleId,
		"accessList": accessList,
	})
}

func (con RoleController) DoAuth(c *gin.Context) {
	// 获取职位id
	roleId, err1 := strconv.Atoi(c.PostForm("role_id"))
	if err1 != nil {
		con.error(c, "参数错误", "/admin/role")
		return
	}
	// 获取权限id  切片
	accessIds := c.PostFormArray("access_node[]")

	// 删除当前职位对应的权限
	roleAccess := database.RoleAccess{}
	database.DB.Where("role_id=?", roleId).Delete(&roleAccess)

	// 增加当前职位对应的权限
	for _, v := range accessIds {
		roleAccess.RoleId = roleId
		accessId, _ := strconv.Atoi(v)
		roleAccess.AccessId = accessId
		database.DB.Create(&roleAccess)
	}
	con.success(c, "授权成功", "/admin/role")
}
