package router

import (
	"final-project/config/enums"
	"final-project/controller"
	"final-project/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/auth/login", controller.Login)
		v1.POST("/auth/register", controller.Register)

		v1.POST("/role", middleware.AuthMiddleware([]string{enums.CREATE, enums.UPDATE}), controller.CreateRole)
		v1.POST("/role-permission", middleware.AuthMiddleware([]string{enums.CREATE, enums.UPDATE}), controller.CreateRolePermission)

		v1.GET("/user", middleware.AuthMiddleware([]string{enums.VIEW}), controller.GetUser)
		v1.PUT("/user", middleware.AuthMiddleware([]string{enums.VIEW}), controller.UpdateUser)
		v1.DELETE("/user/:id", middleware.AuthMiddleware([]string{enums.DELETE}), controller.DeleteUser)
		v1.GET("/user/list", middleware.AuthMiddleware([]string{enums.VIEW}), controller.GetUserList)
		v1.POST("/user-role", middleware.AuthMiddleware([]string{enums.CREATE, enums.UPDATE}), controller.UpdateUserRole)

		v1.POST("/permission", middleware.AuthMiddleware([]string{enums.CREATE, enums.UPDATE}), controller.CreatePermission)

		v1.GET("/dump-user", middleware.AuthMiddleware([]string{enums.CREATE}), controller.DumpUser)
	}

	r.Run(":8080")
}
