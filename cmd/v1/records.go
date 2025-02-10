package v1

import (
	"github.com/gin-gonic/gin"

	"finance_manager/pkg/records"
)

func addRecordRoute(rg *gin.RouterGroup, controller *records.RecordController) {
	recordGroup := rg.Group("/records")
	{
		recordGroup.POST("/create", controller.CreateRecordHandler)
	}
}
