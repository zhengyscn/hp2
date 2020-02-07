package middleware

import (
	"fmt"
	"hp2/utils"
	"net/http"
	"strings"

	"github.com/casbin/casbin"
	"github.com/gin-gonic/gin"
)

// NewAuthorizer returns the authorizer, uses a Casbin enforcer as input
func NewAuthorizer(e *casbin.Enforcer) gin.HandlerFunc {
	a := &BasicAuthorizer{enforcer: e}

	return func(c *gin.Context) {
		if !a.CheckPermission(c.Request) {
			a.RequirePermission(c)
		}
	}
}

// BasicAuthorizer stores the casbin handler
type BasicAuthorizer struct {
	enforcer *casbin.Enforcer
}

// GetUserName gets the user name from the request.
// Currently, only HTTP basic authentication is supported
func (a *BasicAuthorizer) GetUserName(r *http.Request) string {
	username, _, _ := r.BasicAuth()
	return username
}

func (a *BasicAuthorizer) GetUserNameByToken(r *http.Request) string {
	var key = "Token"
	authorizationStr := r.Header.Get("Authorization")
	if len(authorizationStr) == 0 {
		return ""
	}
	authorizationArr := strings.Fields(authorizationStr)
	fmt.Println(authorizationArr)
	if authorizationArr[0] == key && len(authorizationArr) == 2 {
		username, err := utils.ParseToken(authorizationArr[1])
		if err != nil {
			return ""
		}
		return username
	}
	return ""
}

// CheckPermission checks the user/method/path combination from the request.
// Returns true (permission granted) or false (permission forbidden)
func (a *BasicAuthorizer) CheckPermission(r *http.Request) bool {
	//user := a.GetUserName(r)
	user := a.GetUserNameByToken(r)
	method := r.Method
	path := r.URL.Path

	fmt.Printf("user:%s, path:%s, method:%s\n", user, path, method)
	allowed := a.enforcer.Enforce(user, path, method)
	return allowed
}

// RequirePermission returns the 403 Forbidden to the client
func (a *BasicAuthorizer) RequirePermission(c *gin.Context) {
	c.JSON(403, gin.H{
		"code":    "-1",
		"data":    nil,
		"message": "403 Forbidden",
	})
	c.AbortWithStatus(403)
}
