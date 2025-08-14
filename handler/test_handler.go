package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// TestHandler handles requests to the /api/test endpoint
func (h *CashbackReportHandler) TestHandler(c *gin.Context) {
	response := map[string]string{"message": "Test endpoint working"}
	c.JSON(http.StatusOK, response)
}
