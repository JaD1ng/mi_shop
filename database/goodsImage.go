package database

type GoodsImage struct {
	Id      int    `json:"id"`
	GoodsId int    `json:"goods_id"`
	ImgUrl  string `json:"img_url"`
	ColorId int    `json:"color_id"`
	Sort    int    `json:"sort"`
	AddTime int    `json:"add_time"`
	Status  int    `json:"status"`
}

func (goodsImage *GoodsImage) TableName() string {
	return "goods_image"
}
