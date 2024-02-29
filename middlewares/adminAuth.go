package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/ini.v1"
	"mi_shop/database"
	"os"
	"strings"
)

// InitAdminAuth 判断是否登录并验证后台权限
func InitAdminAuth(c *gin.Context) {
	// 获取url访问的地址
	url := c.Request.URL.Path

	// 如果访问的是登录页面则不需要验证
	if url == "/admin/login" || url == "/admin/doLogin" || url == "/admin/captcha" {
		c.Next()
		return
	}

	// 获取session中的用户信息
	session := sessions.Default(c)
	userinfoList := session.Get("userinfo")

	userinfoStr, ok := userinfoList.(string)
	// 没有登录信息则跳转到登录页面
	if !ok {
		c.Redirect(302, "/admin/login")
		return
	}

	// 判断userinfoStr中的信息是否为空
	var userinfoStruct []database.Manager
	err := json.Unmarshal([]byte(userinfoStr), &userinfoStruct)
	if err != nil || len(userinfoStruct) == 0 || userinfoStruct[0].Username == "" {
		c.Redirect(302, "/admin/login")
		return
	}

	// 判断是否有权限访问
	urlPath := strings.Replace(url, "/admin/", "", 1)

	if userinfoStruct[0].IsSuper == 0 && !excludeAuthPath("/"+urlPath) {
		// 1、根据角色获取当前角色的权限列表,然后把权限id放在一个map类型的对象里面
		var roleAccess []database.RoleAccess
		database.DB.Where("role_id=?", userinfoStruct[0].RoleId).Find(&roleAccess)
		roleAccessMap := make(map[int]int)
		for _, v := range roleAccess {
			roleAccessMap[v.AccessId] = v.AccessId
		}

		// 2、获取当前访问的url对应的权限id 判断权限id是否在角色对应的权限
		// url   /admin/manager
		access := database.Access{}
		database.DB.Where("url = ?", urlPath).Find(&access)

		//3、判断当前访问的url对应的权限id 是否在权限列表的id中
		if _, ok := roleAccessMap[access.Id]; !ok {
			c.String(200, "您没有权限")
			c.Abort()
		}
	}
}

// excludeAuthPath 判断是否需要排除权限验证
func excludeAuthPath(urlPath string) bool {
	config, iniErr := ini.Load("./config/app.ini")
	if iniErr != nil {
		fmt.Printf("Fail to read file: %v", iniErr)
		os.Exit(1)
	}

	excludePath := config.Section("").Key("excludeAuthPath").String()
	excludePathSlice := strings.Split(excludePath, ",")

	for _, v := range excludePathSlice {
		if v == urlPath {
			return true
		}
	}
	return false
}
