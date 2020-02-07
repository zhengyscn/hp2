package service

import (
	"hp2/dao"
)

// 用户表验证
type AdminUserRoleDto struct {
	Username string `form:"username" json:"username" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

func (a AdminUserRoleDto) AddUserToRole() bool {
	var err error
	err = dao.Enforce.LoadPolicy()
	ok := dao.Enforce.AddRoleForUser(a.Username, a.Role)
	err = dao.Enforce.SavePolicy()
	if err != nil || !ok {
		return false
	}
	return true
}

type AdminRolePermissionDto struct {
	Role   string `form:"role" json:"role" binding:"required"`
	Uri    string `form:"uri" json:"uri" binding:"required"`
	Method string `form:"method" json:"method" binding:"required"`
}

func (a AdminRolePermissionDto) AddRolePermission() bool {
	var err error
	err = dao.Enforce.LoadPolicy()
	ok := dao.Enforce.AddPermissionForUser(a.Role, a.Uri, a.Method)
	err = dao.Enforce.SavePolicy()
	if err != nil || !ok {
		return false
	}
	return true
}
