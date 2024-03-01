package admin

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

type GoodsCateController struct {
	BaseController
}

func (con GoodsCateController) Index(c *gin.Context) {
	var goodsCateList []database.GoodsCate
	database.DB.Where("pid = 0").Preload("GoodsCateItems").Find(&goodsCateList)

	c.HTML(http.StatusOK, "admin/goodsCate/index.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) Add(c *gin.Context) {
	// 获取顶级分类
	var goodsCateList []database.GoodsCate
	database.DB.Where("pid = 0").Find(&goodsCateList)

	c.HTML(http.StatusOK, "admin/goodsCate/add.html", gin.H{
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoAdd(c *gin.Context) {
	title := c.PostForm("title")
	pid, err1 := strconv.Atoi(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err2 := strconv.Atoi(c.PostForm("sort"))
	status, err3 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err3 != nil {
		con.error(c, "参数错误", "/goodsCate/add")
		return
	}
	if err2 != nil {
		con.error(c, "排序值必须是整数", "/goodsCate/add")
		return
	}

	cateImgDir, _ := util.UploadImg(c, "cate_img")
	goodsCate := database.GoodsCate{
		Title:       title,
		Pid:         pid,
		SubTitle:    subTitle,
		Link:        link,
		Template:    template,
		Keywords:    keywords,
		Description: description,
		CateImg:     cateImgDir,
		Sort:        sort,
		Status:      status,
		AddTime:     int(util.GetUnix()),
	}

	err := database.DB.Create(&goodsCate).Error
	if err != nil {
		con.error(c, "增加数据失败", "/admin/goodsCate/add")
		return
	}
	con.success(c, "增加数据成功", "/admin/goodsCate")
}

func (con GoodsCateController) Edit(c *gin.Context) {
	// 获取要修改的数据
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsCate")
		return
	}

	goodsCate := database.GoodsCate{Id: id}
	database.DB.Find(&goodsCate)

	// 获取顶级分类
	var goodsCateList []database.GoodsCate
	database.DB.Where("pid = 0").Find(&goodsCateList)
	c.HTML(http.StatusOK, "admin/goodsCate/edit.html", gin.H{
		"goodsCate":     goodsCate,
		"goodsCateList": goodsCateList,
	})
}

func (con GoodsCateController) DoEdit(c *gin.Context) {
	id, err1 := strconv.Atoi(c.PostForm("id"))
	title := c.PostForm("title")
	pid, err2 := strconv.Atoi(c.PostForm("pid"))
	link := c.PostForm("link")
	template := c.PostForm("template")
	subTitle := c.PostForm("sub_title")
	keywords := c.PostForm("keywords")
	description := c.PostForm("description")
	sort, err3 := strconv.Atoi(c.PostForm("sort"))
	status, err4 := strconv.Atoi(c.PostForm("status"))

	if err1 != nil || err2 != nil || err4 != nil {
		con.error(c, "参数错误", "/goodsCate/add")
		return
	}
	if err3 != nil {
		con.error(c, "排序值必须是整数", "/goodsCate/add")
		return
	}

	cateImgDir, _ := util.UploadImg(c, "cate_img")

	goodsCate := database.GoodsCate{Id: id}
	database.DB.Find(&goodsCate)
	goodsCate.Title = title
	goodsCate.Pid = pid
	goodsCate.Link = link
	goodsCate.Template = template
	goodsCate.SubTitle = subTitle
	goodsCate.Keywords = keywords
	goodsCate.Description = description
	goodsCate.Sort = sort
	goodsCate.Status = status
	if cateImgDir != "" {
		goodsCate.CateImg = cateImgDir
	}

	err := database.DB.Save(&goodsCate).Error
	if err != nil {
		con.error(c, "修改失败", "/admin/goodsCate/edit?id="+strconv.Itoa(id))
		return
	}
	con.success(c, "修改成功", "/admin/goodsCate")
}

func (con GoodsCateController) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goodsCate")
		return
	}

	// 获取我们要删除的数据
	goodsCate := database.GoodsCate{Id: id}
	database.DB.Find(&goodsCate)
	if goodsCate.Pid == 0 { // 顶级分类
		var goodsCateList []database.GoodsCate
		database.DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)
		if len(goodsCateList) > 0 {
			con.error(c, "当前分类下面有子分类，请删除子分类作以后再来删除这个数据", "/admin/goodsCate")
			return
		}
	}
	// 操作或者菜单
	database.DB.Delete(&goodsCate)
	con.success(c, "删除数据成功", "/admin/goodsCate")
}
