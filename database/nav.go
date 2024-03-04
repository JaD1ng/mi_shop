package database

type Nav struct {
	Id         int
	Title      string
	Link       string
	Position   int
	IsOpennew  int
	Relation   string
	Sort       int
	Status     int
	AddTime    int
	GoodsItems []Goods `gorm:"-"`
}

func (nav *Nav) TableName() string {
	return "nav"
}
