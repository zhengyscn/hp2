package controllers

import (
	"hp2/service"

	"github.com/gin-gonic/gin"
)

type AdminController struct {
	BaseController
}

// 角色中添加用户
func (u *AdminController) AddUserRole(c *gin.Context) {
	var userDto service.AdminUserRoleDto
	if err := c.ShouldBindJSON(&userDto); err == nil {
		if ok := userDto.AddUserToRole(); ok {
			u.success(c, "", "ok")
		} else {
			u.fail(c, "fail")
		}
	} else {
		u.fail(c, err.Error())
	}
}

// 角色中添加权限
func (u *AdminController) AddRolePermission(c *gin.Context) {
	var userDto service.AdminRolePermissionDto
	if err := c.ShouldBindJSON(&userDto); err == nil {
		if ok := userDto.AddRolePermission(); ok {
			u.success(c, "", "ok")
		} else {
			u.fail(c, "fail")
		}
	} else {
		u.fail(c, err.Error())
	}
}
