package users

import (
	"finance_manager/pkg/postgres/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service IUserService
}

func NewUserController(service IUserService) *UserController {
	return &UserController{
		service: service,
	}
}

func (c *UserController) CreateUserHandler(ctx *gin.Context) {
	var params sqlc.CreateUserParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.service.CreateUser(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
