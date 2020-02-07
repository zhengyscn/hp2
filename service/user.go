package service

import (
	"errors"
	"fmt"
	"hp2/dao"
	"hp2/model"

	"github.com/fatih/structs"
)

var daoUser = dao.User{}

// 创建时验证
type UserCreateDto struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (u UserCreateDto) Create(user UserCreateDto) error {
	newUser := model.User{
		UserName: user.UserName,
		Password: user.Password,
	}
	g := daoUser.Create(&newUser)
	if g.Error != nil {
		return g.Error
	}

	if g.RowsAffected >= 1 {
		return nil
	}
	return errors.New(fmt.Sprintf("create user: %s fail", user.UserName))
}

// 登录时验证
type UserAuthDto struct {
	UserName string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (u UserAuthDto) Auth(user UserAuthDto) bool {
	newUser := model.User{
		UserName: user.UserName,
		Password: user.Password,
	}
	return daoUser.CheckByUserPass(&newUser)

}

// 删除时验证
type UserDeleteDto struct {
	ID int `uri:"id" form:"id" json:"id" binding:"required"`
}

func (u UserDeleteDto) Delete(user UserDeleteDto) error {
	newUser := model.User{
		ID: user.ID,
	}
	// 查询
	if ok := daoUser.CheckExistsById(user.ID); !ok {
		return errors.New(fmt.Sprintf("uid: %d not found.", user.ID))
	}

	// 删除
	g := daoUser.DeleteById(&newUser)
	if g.Error != nil {
		return g.Error
	}

	if g.RowsAffected >= 1 {
		return nil
	}
	return errors.New(fmt.Sprintf("delete uid: %d fail", user.ID))
}

// 单条查询时验证
type UserDetailDto struct {
	ID int `uri:"id" form:"id" json:"id" binding:"required"`
}

func (u UserDetailDto) Detail(user UserDetailDto) (map[string]interface{}, error) {
	m := daoUser.GetByUid(user.ID)
	userInfo := structs.Map(&m)
	return userInfo, nil
}

// 单条查询时验证
type UserListDto struct {
}

// 精准查询
func (u UserListDto) List(page, pageSize int, mWhere interface{}) (users []map[string]interface{}, total int) {
	switch where := mWhere.(type) {
	case string:
		userArray, count := daoUser.ListGlobal(page, pageSize, where)
		for _, us := range userArray {
			newUs := structs.Map(us)
			users = append(users, newUs)
		}
		total = count
	case map[string]string:
		userArray, count := daoUser.List(page, pageSize, where)
		for _, us := range userArray {
			newUs := structs.Map(us)
			users = append(users, newUs)
		}
		total = count
	}
	return
}

// 单条查询时验证
type UserUpdateDto struct {
	ID int `uri:"id" form:"id" json:"id"`
}

func (u UserUpdateDto) Update(user map[string]interface{}) (string, error) {
	newUser := model.User{
		ID: u.ID,
	}
	g := daoUser.Update(&newUser, user)
	if g.Error != nil {
		return "UpdateSafe Error", g.Error
	}

	if g.RowsAffected == 1 {
		return "OK", nil
	}
	return "UpdateSafe fail", errors.New("UpdateSafe error")
}
