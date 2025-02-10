package v1

import (
	"finance_manager/cmd/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middleware.LoggerToFile())

	v1 := router.Group("/v1")

	addReportRoute(v1)
	addRecordRoute(v1)
	addImportRoute(v1)

	return router
}
