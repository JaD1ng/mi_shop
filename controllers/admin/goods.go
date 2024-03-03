package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"mi_shop/database"
	"mi_shop/util"
)

var wg sync.WaitGroup

type GoodsController struct {
	BaseController
}

func (con GoodsController) Index(c *gin.Context) {
	var goodsList []database.Goods
	database.DB.Find(&goodsList)

	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{
		"goodsList": goodsList,
	})
}

func (con GoodsController) Add(c *gin.Context) {
	// 获取商品分类
	var goodsCateList []database.GoodsCate
	database.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	// 获取所有颜色信息
	var goodsColorList []database.GoodsColor
	database.DB.Find(&goodsColorList)

	// 获取商品规格包装
	var goodsTypeList []database.GoodsType
	database.DB.Find(&goodsTypeList)

	c.HTML(http.StatusOK, "admin/goods/add.html", gin.H{
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
	})
}

func (con GoodsController) DoAdd(c *gin.Context) {
	// 1、获取表单提交过来的数据 进行判断
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := strconv.Atoi(c.PostForm("cate_id"))
	goodsNumber, _ := strconv.Atoi(c.PostForm("goods_number"))

	// 注意小数点
	marketPrice, _ := util.Float(c.PostForm("market_price"))
	price, _ := util.Float(c.PostForm("price"))

	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	goodsColorArr := c.PostFormArray("goods_color") // 获取的是切片
	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := strconv.Atoi(c.PostForm("is_delete"))
	isHot, _ := strconv.Atoi(c.PostForm("is_hot"))
	isBest, _ := strconv.Atoi(c.PostForm("is_best"))
	isNew, _ := strconv.Atoi(c.PostForm("is_new"))
	goodsTypeId, _ := strconv.Atoi(c.PostForm("goods_type_id"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))
	addTime := int(util.GetUnix())

	// 2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")

	// 3、上传图片   生成缩略图
	goodsImg, _ := util.UploadImg(c, "goods_img")

	// 4、增加商品数据
	goods := database.Goods{
		Title:         title,
		SubTitle:      subTitle,
		GoodsSn:       goodsSn,
		CateId:        cateId,
		ClickCount:    100,
		GoodsNumber:   goodsNumber,
		MarketPrice:   marketPrice,
		Price:         price,
		RelationGoods: relationGoods,
		GoodsAttr:     goodsAttr,
		GoodsVersion:  goodsVersion,
		GoodsGift:     goodsGift,
		GoodsFitting:  goodsFitting,
		GoodsKeywords: goodsKeywords,
		GoodsDesc:     goodsDesc,
		GoodsContent:  goodsContent,
		IsDelete:      isDelete,
		IsHot:         isHot,
		IsBest:        isBest,
		IsNew:         isNew,
		GoodsTypeId:   goodsTypeId,
		Sort:          sort,
		Status:        status,
		AddTime:       addTime,
		GoodsColor:    goodsColorStr,
		GoodsImg:      goodsImg,
	}

	err := database.DB.Create(&goods).Error
	if err != nil {
		con.error(c, "增加商品失败", "/admin/goods/add")
		return
	}

	// 5、增加图库 信息
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := database.GoodsImage{
				GoodsId: goods.Id,
				ImgUrl:  v,
				Sort:    10,
				Status:  1,
				AddTime: int(util.GetUnix()),
			}
			database.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	// 6、增加规格包装
	wg.Add(1)
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")

		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeIdErr := strconv.Atoi(attrIdList[i])
			if attributeIdErr != nil {
				con.error(c, "参数错误", "/admin/goods/add")
				return
			}

			// 获取商品类型属性的数据
			goodsTypeAttributeObj := database.GoodsTypeAttribute{Id: goodsTypeAttributeId}
			database.DB.Find(&goodsTypeAttributeObj)
			// 给商品属性里面增加数据  规格包装
			goodsAttrObj := database.GoodsAttr{
				GoodsId:         goods.Id,
				AttributeTitle:  goodsTypeAttributeObj.Title,
				AttributeType:   goodsTypeAttributeObj.AttrType,
				AttributeId:     goodsTypeAttributeObj.Id,
				AttributeCateId: goodsTypeAttributeObj.CateId,
				AttributeValue:  attrValueList[i],
				Status:          1,
				Sort:            10,
				AddTime:         int(util.GetUnix()),
			}
			database.DB.Create(&goodsAttrObj)
		}
		wg.Done()
	}()
	wg.Wait()
	con.success(c, "增加商品成功", "/admin/goods")
}

func (con GoodsController) Edit(c *gin.Context) {
	// 1、获取要修改的商品数据
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goods")
	}
	goods := database.Goods{Id: id}
	database.DB.Find(&goods)

	// 2、获取商品分类
	var goodsCateList []database.GoodsCate
	database.DB.Where("pid=0").Preload("GoodsCateItems").Find(&goodsCateList)

	// 3、获取所有颜色 以及选中的颜色
	goodsColorSlice := strings.Split(goods.GoodsColor, ",")
	goodsColorMap := make(map[string]string)
	for _, v := range goodsColorSlice {
		goodsColorMap[v] = v
	}

	var goodsColorList []database.GoodsColor
	database.DB.Find(&goodsColorList)
	for i := 0; i < len(goodsColorList); i++ {
		if _, ok := goodsColorMap[strconv.Itoa(goodsColorList[i].Id)]; ok {
			goodsColorList[i].Checked = true
		}
	}

	// 4、商品的图库信息
	var goodsImageList []database.GoodsImage
	database.DB.Where("goods_id=?", goods.Id).Find(&goodsImageList)

	// 5、获取商品类型
	var goodsTypeList []database.GoodsType
	database.DB.Find(&goodsTypeList)

	// 6、获取规格信息
	var goodsAttr []database.GoodsAttr
	database.DB.Where("goods_id=?", goods.Id).Find(&goodsAttr)
	goodsAttrStr := ""

	for _, v := range goodsAttr {
		if v.AttributeType == 1 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: </span> <input type="hidden" name="attr_id_list" value="%v" />   <input type="text" name="attr_value_list" value="%v" /></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else if v.AttributeType == 2 {
			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span><input type="hidden" name="attr_id_list" value="%v" />  <textarea cols="50" rows="3" name="attr_value_list">%v</textarea></li>`, v.AttributeTitle, v.AttributeId, v.AttributeValue)
		} else {
			// 获取当前类型对应的值
			goodsTypeArttribute := database.GoodsTypeAttribute{Id: v.AttributeId}
			database.DB.Find(&goodsTypeArttribute)
			attrValueSlice := strings.Split(goodsTypeArttribute.AttrValue, "\n")

			goodsAttrStr += fmt.Sprintf(`<li><span>%v: 　</span>  <input type="hidden" name="attr_id_list" value="%v" /> `, v.AttributeTitle, v.AttributeId)
			goodsAttrStr += fmt.Sprintf(`<select name="attr_value_list">`)
			for i := 0; i < len(attrValueSlice); i++ {
				if attrValueSlice[i] == v.AttributeValue {
					goodsAttrStr += fmt.Sprintf(`<option value="%v" selected >%v</option>`, attrValueSlice[i], attrValueSlice[i])
				} else {
					goodsAttrStr += fmt.Sprintf(`<option value="%v">%v</option>`, attrValueSlice[i], attrValueSlice[i])
				}
			}
			goodsAttrStr += fmt.Sprintf(`</select>`)
			goodsAttrStr += fmt.Sprintf(`</li>`)
		}
	}

	c.HTML(http.StatusOK, "admin/goods/edit.html", gin.H{
		"goods":          goods,
		"goodsCateList":  goodsCateList,
		"goodsColorList": goodsColorList,
		"goodsTypeList":  goodsTypeList,
		"goodsAttrStr":   goodsAttrStr,
		"goodsImageList": goodsImageList,
	})
}

func (con GoodsController) DoEdit(c *gin.Context) {
	// 1、获取表单提交过来的数据
	id, err := strconv.Atoi(c.PostForm("id"))
	if err != nil {
		con.error(c, "参数错误", "/admin/goods")
		return
	}
	title := c.PostForm("title")
	subTitle := c.PostForm("sub_title")
	goodsSn := c.PostForm("goods_sn")
	cateId, _ := strconv.Atoi(c.PostForm("cate_id"))
	goodsNumber, _ := strconv.Atoi(c.PostForm("goods_number"))
	marketPrice, _ := util.Float(c.PostForm("market_price")) // 注意小数点
	price, _ := util.Float(c.PostForm("price"))
	relationGoods := c.PostForm("relation_goods")
	goodsAttr := c.PostForm("goods_attr")
	goodsVersion := c.PostForm("goods_version")
	goodsGift := c.PostForm("goods_gift")
	goodsFitting := c.PostForm("goods_fitting")
	goodsColorArr := c.PostFormArray("goods_color") // 获取的是切片
	goodsKeywords := c.PostForm("goods_keywords")
	goodsDesc := c.PostForm("goods_desc")
	goodsContent := c.PostForm("goods_content")
	isDelete, _ := strconv.Atoi(c.PostForm("is_delete"))
	isHot, _ := strconv.Atoi(c.PostForm("is_hot"))
	isBest, _ := strconv.Atoi(c.PostForm("is_best"))
	isNew, _ := strconv.Atoi(c.PostForm("is_new"))
	goodsTypeId, _ := strconv.Atoi(c.PostForm("goods_type_id"))
	sort, _ := strconv.Atoi(c.PostForm("sort"))
	status, _ := strconv.Atoi(c.PostForm("status"))

	// 2、获取颜色信息 把颜色转化成字符串
	goodsColorStr := strings.Join(goodsColorArr, ",")

	// 3、修改数据
	goods := database.Goods{Id: id}
	database.DB.Find(&goods)
	goods.Title = title
	goods.SubTitle = subTitle
	goods.GoodsSn = goodsSn
	goods.CateId = cateId
	goods.GoodsNumber = goodsNumber
	goods.MarketPrice = marketPrice
	goods.Price = price
	goods.RelationGoods = relationGoods
	goods.GoodsAttr = goodsAttr
	goods.GoodsVersion = goodsVersion
	goods.GoodsGift = goodsGift
	goods.GoodsFitting = goodsFitting
	goods.GoodsKeywords = goodsKeywords
	goods.GoodsDesc = goodsDesc
	goods.GoodsContent = goodsContent
	goods.IsDelete = isDelete
	goods.IsHot = isHot
	goods.IsBest = isBest
	goods.IsNew = isNew
	goods.GoodsTypeId = goodsTypeId
	goods.Sort = sort
	goods.Status = status
	goods.GoodsColor = goodsColorStr

	// 4、上传图片   生成缩略图
	goodsImg, err := util.UploadImg(c, "goods_img")
	if err == nil && len(goodsImg) > 0 {
		goods.GoodsImg = goodsImg
	}

	err = database.DB.Save(&goods).Error
	if err != nil {
		con.error(c, "修改失败", "/admin/goods/edit?id="+strconv.Itoa(id))
		return
	}

	// 5、修改图库 增加图库信息
	wg.Add(1)
	go func() {
		goodsImageList := c.PostFormArray("goods_image_list")
		for _, v := range goodsImageList {
			goodsImgObj := database.GoodsImage{
				GoodsId: goods.Id,
				ImgUrl:  v,
				Sort:    10,
				Status:  1,
				AddTime: int(util.GetUnix()),
			}
			database.DB.Create(&goodsImgObj)
		}
		wg.Done()
	}()

	// 6、修改规格包装  1、删除当前商品下面的规格包装   2、重新执行增加
	// 6.1删除当前商品下面的规格包装
	goodsAttrObj := database.GoodsAttr{}
	database.DB.Where("goods_id=?", goods.Id).Delete(&goodsAttrObj)
	// 6.2、重新执行增加
	wg.Add(1)
	go func() {
		attrIdList := c.PostFormArray("attr_id_list")
		attrValueList := c.PostFormArray("attr_value_list")
		for i := 0; i < len(attrIdList); i++ {
			goodsTypeAttributeId, attributeIdErr := strconv.Atoi(attrIdList[i])
			if attributeIdErr == nil {
				// 获取商品类型属性的数据
				goodsTypeAttributeObj := database.GoodsTypeAttribute{Id: goodsTypeAttributeId}
				database.DB.Find(&goodsTypeAttributeObj)
				// 给商品属性里面增加数据  规格包装
				goodsAttrObj := database.GoodsAttr{
					GoodsId:         goods.Id,
					AttributeTitle:  goodsTypeAttributeObj.Title,
					AttributeType:   goodsTypeAttributeObj.AttrType,
					AttributeId:     goodsTypeAttributeObj.Id,
					AttributeCateId: goodsTypeAttributeObj.CateId,
					AttributeValue:  attrValueList[i],
					Status:          1,
					Sort:            10,
					AddTime:         int(util.GetUnix()),
				}
				database.DB.Create(&goodsAttrObj)
			}
		}
		wg.Done()
	}()
	wg.Wait()
	con.success(c, "修改数据成功", "/admin/goods")
}

// GoodsTypeAttribute 获取并返回商品类型属性
func (con GoodsController) GoodsTypeAttribute(c *gin.Context) {
	cateId, err1 := strconv.Atoi(c.Query("cateId"))
	var goodsTypeAttributeList []database.GoodsTypeAttribute
	err2 := database.DB.Where("cate_id = ?", cateId).Find(&goodsTypeAttributeList).Error
	if err1 != nil || err2 != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"result":  "",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"result":  goodsTypeAttributeList,
	})
}

// ImageUpload 图片上传
func (con GoodsController) ImageUpload(c *gin.Context) {
	imgDir, err := util.UploadImg(c, "file") // 注意：可以在网络里面看到传递的参数
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"link": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"link": "/" + imgDir,
	})
}
