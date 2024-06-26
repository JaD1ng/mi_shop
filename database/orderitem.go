package database

type OrderItem struct {
	Id           int
	OrderId      int
	Uid          int
	ProductTitle string
	ProductId    int
	ProductImg   string
	ProductPrice float64
	ProductNum   int
	GoodsVersion string
	GoodsColor   string
	AddTime      int
}

func (orderItem *OrderItem) TableName() string {
	return "order_item"
}
