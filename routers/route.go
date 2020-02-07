package routers

import (
	"hp2/config"
	"hp2/controllers"
	"hp2/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	gin.SetMode(config.ServerConfig.RunMode)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	apiV1 := r.Group("/api/auth")
	tokenController := &controllers.TokenController{}
	{
		apiV1.POST("/", tokenController.Generator)
		apiV1.GET("/", tokenController.Parse)
	}

	apiV2 := r.Group("/api/v1")
	apiV2.Use(middleware.NewAuthorizer(middleware.RestfulPermission())) // auth permission
	userController := &controllers.UserController{}
	{
		apiV2.GET("/users", userController.List)
		apiV2.GET("/users/:id", userController.Detail)
		apiV2.POST("/users", userController.Create)
		apiV2.PUT("/users/:id", userController.Update)
		apiV2.DELETE("/users/:id", userController.Delete)

	}

	apiV3 := r.Group("/api/admin")
	adminController := &controllers.AdminController{}
	{
		apiV3.POST("/userRole", adminController.AddUserRole)
		apiV3.POST("/rolePermission", adminController.AddRolePermission)
	}

	return r
}
