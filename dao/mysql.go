package dao

import (
	"fmt"
	"hp2/config"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	dbMaster *gorm.DB
	dbSlave  *gorm.DB
)

// 如果建立连接失败，退出程序.
func Setup() {
	dbMaster = mustBuildConnect(generataDsn("master"))
	fmt.Println("Master connection mysql success.")
	if config.SlaveDatabaseConfig.Enable {
		dbSlave = mustBuildConnect(generataDsn("slave"))
		fmt.Println("Slave connection mysql success.")
	}
}

func generataDsn(dbtype string) string {
	switch dbtype {
	case "master":
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.MasterDatabaseConfig.User,
			config.MasterDatabaseConfig.Password,
			config.MasterDatabaseConfig.Host,
			config.MasterDatabaseConfig.Port,
			config.MasterDatabaseConfig.Name)
	case "slave":
		return fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			config.SlaveDatabaseConfig.User,
			config.SlaveDatabaseConfig.Password,
			config.SlaveDatabaseConfig.Host,
			config.SlaveDatabaseConfig.Port,
			config.SlaveDatabaseConfig.Name)
	default:
		return ""
	}
}

func mustBuildConnect(dsn string) (db *gorm.DB) {
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(fmt.Sprintf("Connect MySQL failed, Plase check: %v\n", err.Error()))
	}

	db.DB().SetConnMaxLifetime(100 * time.Second) // 最大连接周期，超过时间的连接就close
	db.DB().SetMaxOpenConns(100)                  // 设置最大连接数
	db.DB().SetMaxIdleConns(16)                   // 设置闲置连接数

	db.LogMode(true)

	// 禁用表名的复数形式
	db.SingularTable(true)
	return
}

func CloseMasterDB() error {
	return dbMaster.Close()
}

func CloseSlaveDB() error {
	return dbSlave.Close()
}
func Using(dbName string) *gorm.DB {
	switch dbName {
	case "":
		return dbMaster
	case "master":
		return dbMaster
	case "slave":
		return dbSlave
	default:
		return nil
	}
}
