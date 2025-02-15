package users

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, controller *UserController) {
	recordGroup := rg.Group("/users")
	{
		recordGroup.POST("/signup", controller.CreateUserHandler)
	}
}
