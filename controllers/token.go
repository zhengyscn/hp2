package controllers

import (
	"fmt"
	"hp2/service"
	"hp2/utils"

	"github.com/gin-gonic/gin"
)

type TokenController struct {
	BaseController
}

// 生成Token
func (u *TokenController) Generator(c *gin.Context) {
	var userDto service.UserAuthDto
	if err := c.ShouldBindJSON(&userDto); err == nil {
		fmt.Printf("userDto %#v\n", userDto)
		if ok := userDto.Auth(userDto); !ok {
			u.fail(c, "auth failed.")
			return
		}
		token, err := utils.GenerateToken(userDto.UserName)
		if err != nil {
			u.fail(c, err.Error())
		} else {
			u.success(c, token, "")
		}
	} else {
		u.fail(c, err.Error())
	}
}

// 解析Token
func (u *TokenController) Parse(c *gin.Context) {
	var userDto service.UserTokenDto
	if err := c.ShouldBindQuery(&userDto); err == nil {
		fmt.Printf("token: %#v\n", userDto)
		username, err := utils.ParseToken(userDto.Token)
		if err != nil {
			u.fail(c, err.Error())
		} else {
			u.success(c, username, "")
		}
	} else {
		u.fail(c, err.Error())
	}
}
