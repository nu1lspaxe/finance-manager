package records

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(rg *gin.RouterGroup, controller *RecordController) {
	recordGroup := rg.Group("/records")
	{
		recordGroup.POST("/create", controller.CreateRecordHandler)
	}
}
