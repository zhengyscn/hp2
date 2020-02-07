package dao

import (
	"fmt"
	"hp2/model"
)

// 同步[批量]创建表
func AutoMigrate() {
	fmt.Println(dbMaster)
	dbMaster.Set("gorm:table_options",
		"ENGINE=InnoDB DEFAULT CHARSET=utf8").AutoMigrate(
		&model.User{},
		&model.Method{},
		&model.Permission{},
	)
}
