package service

// 用户表验证
type UserTokenDto struct {
	Token string `form:"token" json:"token" binding:"required"`
}
