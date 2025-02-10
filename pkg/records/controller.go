package records

import (
	"finance_manager/pkg/postgres/sqlc"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecordController struct {
	service IRecordService
}

func NewRecordController(service IRecordService) *RecordController {
	return &RecordController{
		service: service,
	}
}

func (c *RecordController) CreateRecordHandler(ctx *gin.Context) {
	var params sqlc.CreateRecordParams

	if err := ctx.ShouldBindJSON(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := c.service.CreateRecord(ctx, params)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Record created successfully"})
}
