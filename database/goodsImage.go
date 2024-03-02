package database

type GoodsImage struct {
	Id      int
	GoodsId int
	ImgUrl  string
	ColorId int
	Sort    int
	AddTime int
	Status  int
}

func (goodsImage *GoodsImage) TableName() string {
	return "goods_image"
}
