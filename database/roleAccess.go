package database

type RoleAccess struct {
	AccessId int
	RoleId   int
}

func (roleAccess *RoleAccess) TableName() string {
	return "role_access"
}
