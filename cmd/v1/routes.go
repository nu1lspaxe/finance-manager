package v1

import (
	"finance_manager/cmd/middleware"
	"finance_manager/pkg/records"
	"finance_manager/pkg/users"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitializeRouter(pool *pgxpool.Pool) *gin.Engine {

	iRecordService := records.NewRecordService(pool)
	recordController := records.NewRecordController(iRecordService)

	iUserService := users.NewUserService(pool)
	userController := users.NewUserController(iUserService)

	router := SetupRouter(recordController, userController)
	return router
}

func SetupRouter(
	recordController *records.RecordController,
	userController *users.UserController,
) *gin.Engine {

	router := gin.Default()

	router.Use(middleware.LoggerToFile())

	v1 := router.Group("/v1")

	records.RegisterRoutes(v1, recordController)
	users.RegisterRoutes(v1, userController)

	return router
}
