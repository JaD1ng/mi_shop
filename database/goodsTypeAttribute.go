package database

type GoodsTypeAttribute struct {
	Id        int
	CateId    int
	Title     string
	AttrType  int
	AttrValue string
	Status    int
	Sort      int
	AddTime   int
}

func (goodsTypeAttribute *GoodsTypeAttribute) TableName() string {
	return "goods_type_attribute"
}
