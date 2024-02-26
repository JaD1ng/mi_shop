package database

type Manager struct {
	Id       int
	Username string
	Password string
	Mobile   string
	Email    string
	Status   int
	RoleId   int
	AddTime  int
	IsSuper  int
}

// TableName 表示配置操作数据库的表名称
func (m *Manager) TableName() string {
	return "manager"
}
