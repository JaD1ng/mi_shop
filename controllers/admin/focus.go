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
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/focus")
		return
	}

	focus := database.Focus{Id: id}
	database.DB.Find(&focus)
	c.HTML(http.StatusOK, "admin/focus/edit.html", gin.H{
		"focus": focus,
	})
}

func (con FocusController) DoEdit(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	focusType, err2 := strconv.Atoi(c.PostForm("focus_type"))
	link := c.PostForm("link")
	sort, err3 := strconv.Atoi(c.PostForm("sort"))
	status, err4 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.error(c, "非法请求", "/admin/focus")
	}
	if err3 != nil {
		con.error(c, "请输入正确的排序值", "/admin/focus/edit?id="+strconv.Itoa(id))
	}

	// 上传文件
	focusImg, _ := util.UploadImg(c, "focus_img")

	focus := database.Focus{Id: id}
	database.DB.Find(&focus)
	focus.Title = title
	focus.FocusType = focusType
	focus.Link = link
	focus.Sort = sort
	focus.Status = status
	if focusImg != "" {
		focus.FocusImg = focusImg
	}

	err := database.DB.Save(&focus).Error
	if err != nil {
		con.error(c, "修改数据失败请重新尝试", "/admin/focus/edit?id="+strconv.Itoa(id))
		return
	}
	con.success(c, "增加轮播图成功", "/admin/focus")
}

func (con FocusController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/focus")
		return
	}

	focus := database.Focus{Id: id}
	database.DB.Delete(&focus)
	// 根据自己的需要 要不要删除图片
	// os.Remove("static/upload/20210915/1631694117.jpg")
	con.success(c, "删除数据成功", "/admin/focus")
}
