package database

type User struct {
	Id       int
	Username string
	Age      int
	Email    string
	AddTime  int
}

// TableName 表示配置操作数据库的表名称
func (u *User) TableName() string {
	return "user"
}
