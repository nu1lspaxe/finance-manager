package v1

import (
	"finance_manager/pkg/users"

	"github.com/gin-gonic/gin"
)

func addUserRoute(rg *gin.RouterGroup, controller *users.UserController) {
	recordGroup := rg.Group("/users")
	{
		recordGroup.POST("/signup", controller.CreateUserHandler)
	}
}
