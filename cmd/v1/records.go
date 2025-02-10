package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addRecordRoute(rg *gin.RouterGroup) {
	recordGroup := rg.Group("/records")
	{
		recordGroup.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{"message": "Record created"})
		})
	}
}
