package controllers

import (
	"fmt"
	"hp2/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

// /api/v1/users/<uid>
func (t *UserController) Detail(c *gin.Context) {
	var userDto service.UserDetailDto
	if err := c.ShouldBindUri(&userDto); err == nil {
		fmt.Printf("%#v\n", userDto)
		userInfo, err := userDto.Detail(userDto)
		if err != nil {
			t.fail(c, err.Error())
		} else {
			t.success(c, userInfo, "OK")
		}
	} else {
		t.fail(c, err.Error())
	}
}

// /api/v1/users?page=1&page_size=15
// /api/v1/users?page=1&page_size=15&default=mk
// /api/v1/users?page=1&page_size=15&username=mk&password=123
func (t *UserController) List(c *gin.Context) {
	var userDto service.UserListDto
	page, pageSize, mWhere := t.GetPaginator(c)
	users, total := userDto.List(page, pageSize, mWhere)
	t.successPaginator(c, users, total)
}

func (t *UserController) Create(c *gin.Context) {
	var userDto service.UserCreateDto
	if err := c.ShouldBindJSON(&userDto); err == nil {
		fmt.Printf("%#v\n", userDto)
		err := userDto.Create(userDto)
		if err != nil {
			t.fail(c, err.Error())
		} else {
			t.success(c, nil, "OK")
		}
	} else {
		t.fail(c, err.Error())
	}
}

func (t *UserController) Delete(c *gin.Context) {
	var userDto service.UserDeleteDto
	if err := c.ShouldBindUri(&userDto); err == nil {
		fmt.Printf("%#v\n", userDto)
		err := userDto.Delete(userDto)
		if err != nil {
			t.fail(c, err.Error())
		} else {
			t.success(c, nil, "OK")
		}
	} else {
		t.fail(c, err.Error())
	}
}

func (t *UserController) Update(c *gin.Context) {
	var (
		err     error
		userDto service.UserUpdateDto
		dataMap map[string]interface{}
	)

	if dataMap, err = t.GetBodyData(c); err != nil {
		t.fail(c, "GetBodyData fail.")
	} else {
		if err := c.ShouldBindUri(&userDto); err != nil {
			t.fail(c, err.Error())
		} else {
			resp, err := userDto.Update(dataMap)
			if err != nil {
				t.fail(c, resp)
			} else {
				t.success(c, "", "OK")
			}
		}

	}
}
