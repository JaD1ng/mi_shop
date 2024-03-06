package admin

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"mi_shop/database"
	"mi_shop/util"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(c *gin.Context) {
	// 获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")

	// 类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)
	if !ok {
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	// 1、获取用户信息
	var userinfoStruct []database.Manager
	json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

	// 2、获取所有的权限
	var accessList []database.Access
	database.DB.Where("module_id=?", 0).Preload("AccessItem", func(db *gorm.DB) *gorm.DB {
		return db.Order("access.sort DESC")
	}).Order("sort DESC").Find(&accessList)

	// 3、获取当前职位拥有的权限 ，并把权限id放在一个map对象里面
	var roleAccess []database.RoleAccess
	database.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
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

	c.HTML(http.StatusOK, "admin/main/index.html", gin.H{
		"username":   userinfoStruct[0].Username,
		"accessList": accessList,
		"isSuper":    userinfoStruct[0].IsSuper,
	})
}

func (con MainController) Welcome(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/main/welcome.html", gin.H{})
}

// ChangeStatus 修改状态
func (con MainController) ChangeStatus(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")

	err = database.DB.Exec("update "+table+" set "+field+"=ABS("+field+"-1) where id=?", id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改失败，请重试",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

func (con MainController) ChangeNum(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "参数错误",
		})
		return
	}

	table := c.Query("table")
	field := c.Query("field")
	num := c.Query("num")

	err = database.DB.Exec("update "+table+" set "+field+"="+num+" where id=?", id).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "修改数据失败",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "修改成功",
	})
}

func (con MainController) FlushAll(c *gin.Context) {
	util.CacheDb.FlushAll()
	con.success(c, "清除redis缓存数据成功", "/admin")
}
