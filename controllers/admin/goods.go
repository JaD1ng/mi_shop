package admin

import (
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
	c.HTML(http.StatusOK, "admin/goods/index.html", gin.H{})
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
