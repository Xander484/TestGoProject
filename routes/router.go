package routes

import (
	"database/sql"

	"myapp/handler"
	"myapp/service"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Handlers that need DB
	service := &service.CashbackReportService{DB: db}
	handler := &handler.CashbackReportHandler{Service: service}

	api := r.Group("/api")
	{
		api.GET("/test", handler.TestHandler)
		api.GET("/user", handler.UserHandler)
		api.GET("/cashback", handler.CashbackHistory) // now works with Gin
	}

	return r
}
