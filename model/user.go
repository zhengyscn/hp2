package model

// 用户表
type User struct {
	ID       int    `gorm:"primary_key" json:"id"`
	UserName string `gorm:"unique_index" json:"username"`
	Password string `json:"password"`
}

// 模糊查询的字段
func (u User) DefaultSearchFields() []string {
	return []string{"user_name", "password"}
}

func (u User) Display() []string {
	return []string{"user_name", "password"}
}
