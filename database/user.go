package database

type User struct {
	Id       int
	Phone    string
	Password string
	AddTime  int
	LastIp   string
	Email    string
	Status   int
}

// TableName 表示配置操作数据库的表名称
func (u *User) TableName() string {
	return "user"
}
