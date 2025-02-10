package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addReportRoute(rg *gin.RouterGroup) {
	reportGroup := rg.Group("/reports")
	{
		reportGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Report generated"})
		})
	}
}
