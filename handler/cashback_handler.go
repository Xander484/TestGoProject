package handler

import (
	"myapp/models"
	"myapp/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CashbackReportHandler struct {
	Service *service.CashbackReportService
}

func (h *CashbackReportHandler) CashbackHistory(c *gin.Context) {
	var filters models.CashbackHistoryFilters
	if err := c.ShouldBind(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.Service.GetCashbackHistory(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
