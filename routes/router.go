package routes

import (
	"database/sql"

	"myapp/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	// Handlers that need DB
	cashbackHandler := handler.CashbackReportHandler{DB: db}

	api := r.Group("/api")
	{
		api.GET("/test", handler.TestHandler)
		api.GET("/user", handler.UserHandler)
		api.GET("/cashback", cashbackHandler.CashbackHistory) // now works with Gin
	}

	return r
}
