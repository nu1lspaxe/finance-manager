package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addImportRoute(rg *gin.RouterGroup) {
	importGroup := rg.Group("/imports")
	{
		importGroup.POST("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "Transactions imported and reconciled"})
		})
	}
}
