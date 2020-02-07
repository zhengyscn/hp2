package dao

import (
	"fmt"
	"hp2/model"

	"github.com/jinzhu/gorm"
)

type User struct {
}

// Create - new user
func (u User) Create(user *model.User) *gorm.DB {
	db := Using("")
	return db.Create(user)
}

// Update - update user
func (u User) Update(user *model.User, ups map[string]interface{}) *gorm.DB {
	db := Using("")
	return db.Model(user).Update(ups)
}

// Update - update user
func (u User) UpdateSafe(user *model.User) *gorm.DB {
	db := Using("")
	m := model.User{}
	return db.Model(&m).Updates(user)
}

// List - users list
func (u User) List(page, pageSize int, mWhere map[string]string) ([]model.User, int) {
	var total int
	var sqlRawStr string
	var users []model.User
	var sqlReplaceSlice []interface{}

	m := model.User{}

	// 定义字段名称
	idx := 0
	for key, value := range mWhere {
		where := fmt.Sprintf("%v = ?", key)
		if idx >= 1 {
			sqlRawStr += " and " + where
		} else {
			sqlRawStr += where
		}
		sqlReplaceSlice = append(sqlReplaceSlice, value)
		idx++
	}

	// db操作
	db := Using("")

	// 获取分页数据
	g := db.Model(&m).Where(sqlRawStr, sqlReplaceSlice...)
	g.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)

	// 获取总条数
	g.Count(&total)

	return users, total
}

// List - users list
func (u User) ListGlobal(page, pageSize int, mWhere string) ([]model.User, int) {
	var total int
	var sqlRawStr string
	var users []model.User
	var sqlReplaceSlice []interface{}

	m := model.User{}

	// 定义字段名称
	if len(mWhere) != 0 {
		fields := m.DefaultSearchFields()
		for idx, field := range fields {
			where := fmt.Sprintf("%v like ?", field)
			if idx >= 1 {
				sqlRawStr += " or " + where
			} else {
				sqlRawStr += where
			}
			sqlReplaceSlice = append(sqlReplaceSlice, fmt.Sprintf("%%%s%%", mWhere))
		}

		// db操作
		db := Using("")

		// 获取分页数据
		g := db.Model(&m).Where(sqlRawStr, sqlReplaceSlice...)
		g.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)

		// 获取总条数
		g.Count(&total)
	} else {
		// db操作
		db := Using("")

		// 获取分页数据
		g := db.Model(&m)
		g.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users)

		// 获取总条数
		g.Count(&total)
	}

	return users, total
}

// Update - update user
func (u User) UpdateSafeSec(ups map[string]interface{}) *gorm.DB {
	db := Using("")
	m := model.User{}
	return db.Model(&m).Omit("username").Updates(ups)
}

// Delete - delete user by id
func (u User) DeleteById(user *model.User) *gorm.DB {
	db := Using("")
	return db.Delete(user)

}

// CheckByUserPass ...
func (u User) CheckByUserPass(user *model.User) bool {
	var c int
	db := Using("")
	m := model.User{}
	db.Where("user_name = ? and password = ?", user.UserName, user.Password).
		First(&m).Count(&c)
	if c == 1 {
		return true
	}
	return false

}

// GetByUserName - get user from name
func (u User) GetByUserName(username string) model.User {
	db := Using("")
	m := model.User{}
	db.Where("user_name = ?", username).First(&m)
	return m
}

// GetByUserName - get user from name
func (u User) GetByUid(uid int) model.User {
	db := Using("")
	m := model.User{}
	db.Where("id = ?", uid).First(&m)
	return m
}

// CheckUsernameExists - find username
func (u User) CheckExistsByUsername(username string) bool {
	var c int
	m := model.User{}
	db := Using("")
	db.Where("user_name = ?", username).First(&m).Count(&c)
	if c >= 1 {
		return true
	}
	return false
}

// CheckUsernameExists - find username
func (u User) CheckExistsById(uid int) bool {
	var c int
	m := model.User{}
	db := Using("")
	db.Where("id = ?", uid).First(&m).Count(&c)
	if c == 1 {
		return true
	}
	return false
}
