package admin

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type FocusController struct {
	BaseController
}

func (con FocusController) Index(c *gin.Context) {
	var focusList []database.Focus
	database.DB.Find(&focusList)
	c.HTML(http.StatusOK, "admin/focus/index.html", gin.H{
		"focusList": focusList,
	})
}

func (con FocusController) Add(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/add.html", gin.H{})
}

func (con FocusController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	link := c.PostForm("link")
	focusType, err := strconv.Atoi(c.PostForm("focus_type"))
	if err != nil {
		con.error(c, "非法请求", "/admin/focus/add")
		return
	}

	sort, err := strconv.Atoi(c.PostForm("sort"))
	if err != nil {
		con.error(c, "请输入正确的排序值", "/admin/focus/add")
		return
	}

	status, err := strconv.Atoi(c.PostForm("status"))
	if err != nil {
		con.error(c, "非法请求", "/admin/focus/add")
		return
	}

	// 上传文件
	focusImg, err := util.UploadImg(c, "focus_img")
	if err != nil {
		fmt.Println(err)
	}

	focus := database.Focus{
		Title:     title,
		FocusType: focusType,
		FocusImg:  focusImg,
		Link:      link,
		Sort:      sort,
		Status:    status,
		AddTime:   int(util.GetUnix()),
	}

	err = database.DB.Create(&focus).Error
	if err != nil {
		con.error(c, "增加轮播图失败", "/admin/focus/add")
		return
	}
	con.success(c, "增加轮播图成功", "/admin/focus")
}

func (con FocusController) Edit(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{})
}

func (con FocusController) DoEdit(c *gin.Context) {
	c.String(http.StatusOK, "执行编辑")

}

func (con FocusController) Delete(c *gin.Context) {
	c.String(http.StatusOK, "执行删除")
}
