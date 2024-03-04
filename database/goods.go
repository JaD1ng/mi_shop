package database

type Goods struct {
	Id            int
	Title         string
	SubTitle      string
	GoodsSn       string
	CateId        int
	ClickCount    int
	GoodsNumber   int
	Price         float64
	MarketPrice   float64
	RelationGoods string
	GoodsAttr     string
	GoodsVersion  string
	GoodsImg      string
	GoodsGift     string
	GoodsFitting  string
	GoodsColor    string
	GoodsKeywords string
	GoodsDesc     string
	GoodsContent  string
	IsDelete      int
	IsHot         int
	IsBest        int
	IsNew         int
	GoodsTypeId   int
	Sort          int
	Status        int
	AddTime       int
}

func (goods *Goods) TableName() string {
	return "goods"
}

// GetGoodsByCategory 根据商品分类获取推荐商品
func GetGoodsByCategory(cateId int, goodsType string, limitNum int) (goodsList []Goods) {
	// 判断cateId 是否是顶级分类
	goodsCate := GoodsCate{Id: cateId}
	DB.Find(&goodsCate)
	var tempSlice []int
	if goodsCate.Pid == 0 { // 顶级分类
		// 获取顶级分类下面的二级分类
		var goodsCateList []GoodsCate
		DB.Where("pid = ?", goodsCate.Id).Find(&goodsCateList)

		for i := 0; i < len(goodsCateList); i++ {
			tempSlice = append(tempSlice, goodsCateList[i].Id)
		}
	}
	tempSlice = append(tempSlice, cateId)

	// 获取商品
	where := "cate_id in ?"
	switch goodsType {
	case "hot":
		where += " AND is_hot=1"
		break
	case "best":
		where += " AND is_best=1"
		break
	case "new":
		where += " AND is_new=1"
		break
	default:
		break
	}

	DB.Where(where, tempSlice).Select("id,title,price,goods_img,sub_title").Limit(limitNum).Order("sort desc").Find(&goodsList)
	return
}
