package middleware

import (
	"hp2/dao"

	"github.com/casbin/casbin"
)

func RestfulPermission() *casbin.Enforcer {
	return dao.Enforce
}
