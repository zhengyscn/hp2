package model

type Method struct {
	ID   int    `gorm:"primary_key" json:"id"`
	Name string `gorm:"unique_index" json:"name"`
}

// 权限
type Permission struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Name     string `gorm:"unique_index" json:"name"`
	MethodId uint   `json:"method_id"`
	Methods  Method `gorm:"ForeignKey:MethodId;AssociationForeignKey:ID"`
	Remark   string `json:"remark"`
}

/*
// 组
type Role struct {
	ID          int          `gorm:"primary_key" json:"id"`
	Name        string       `gorm:"unique_index" json:"name"`
	Users       []User       `gorm:"many2many:role_users"`
	Permissions []Permission `gorm:"many2many:role_users"`
}
*/
