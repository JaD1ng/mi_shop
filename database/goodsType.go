package database

type GoodsType struct {
	Id          int
	Title       string
	Description string
	Status      int
	AddTime     int
}

func (goodsType *GoodsType) TableName() string {
	return "goods_type"
}
