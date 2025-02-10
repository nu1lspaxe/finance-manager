package v1

import (
	"finance_manager/cmd/middleware"
	"finance_manager/pkg/postgres/sqlc"
	"finance_manager/pkg/records"
	"finance_manager/pkg/users"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

func SetupRouter(conn *pgx.Conn) *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerToFile())

	v1 := router.Group("/v1")

	recordService := records.NewRecordService(sqlc.New(conn))
	recordController := records.NewRecordController(recordService)
	addRecordRoute(v1, recordController)

	userService := users.NewUserService(sqlc.New(conn))
	userController := users.NewUserController(userService)
	addUserRoute(v1, userController)

	return router
}
