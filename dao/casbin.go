package dao

import (
	"fmt"
	"hp2/config"
	"path/filepath"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
)

var (
	Enforce *casbin.Enforcer
)

func NewCasbin() error {
	dir, _ := filepath.Abs(filepath.Dir("."))
	rbacPath := fmt.Sprintf("%s/config/%s", dir, config.CasbinDatabaseConfig.RbacConfigPath)
	fmt.Printf("rbacPath: %s\n", rbacPath)
	//policyPath string = config.CasbinDatabaseConfig.PolicyPath
	//var dsn string = "moyu:12345678@tcp(192.168.3.100:3306)/"
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/",
		config.MasterDatabaseConfig.User,
		config.MasterDatabaseConfig.Password,
		config.MasterDatabaseConfig.Host,
		config.MasterDatabaseConfig.Port,
	)
	a := gormadapter.NewAdapter("mysql", dsn)
	Enforce = casbin.NewEnforcer(rbacPath, a)
	fmt.Println("Enforce", Enforce)
	err := Enforce.LoadPolicy()
	if err != nil {
		return err
	}
	return nil
}
