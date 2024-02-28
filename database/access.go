package database

type Access struct {
	Id          int
	ModuleName  string
	ActionName  string
	Type        int // 1:模块, 2:菜单, 3:操作
	Url         string
	ModuleId    int
	Sort        int
	Description string
	Status      int
	AddTime     int
	AccessItem  []Access `gorm:"foreignKey:ModuleId;references:Id"`
}

func (access *Access) TableName() string {
	return "access"
}
