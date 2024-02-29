package admin

import (
	"encoding/json"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"net/http"
)

type MainController struct {
	BaseController
}

func (con MainController) Index(c *gin.Context) {
	//获取userinfo 对应的session
	session := sessions.Default(c)
	userinfo := session.Get("userinfo")

	//类型断言 来判断 userinfo是不是一个string
	userinfoStr, ok := userinfo.(string)
	if !ok {
		c.Redirect(http.StatusFound, "/admin/login")
		return
	}

	//1、获取用户信息
	var userinfoStruct []database.Manager
	json.Unmarshal([]byte(userinfoStr), &userinfoStruct)

	//2、获取所有的权限
	var accessList []database.Access
	database.DB.Where("module_id=?", 0).Preload("AccessItem").Find(&accessList)

	//3、获取当前角色拥有的权限 ，并把权限id放在一个map对象里面
	var roleAccess []database.RoleAccess
	database.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
	roleAccessMap := make(map[int]int)
	for _, v := range roleAccess {
		roleAccessMap[v.AccessId] = v.AccessId
	}

	//4、循环遍历所有的权限数据，判断当前权限的id是否在角色权限的Map对象中,如果是的话给当前数据加入checked属性
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
