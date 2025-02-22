package routes

import (
	"github.com/adedejiosvaldo/safe_url/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/scan-url", handlers.HandleScanRequest)
	router.GET("/ping", handlers.HealthCheck)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/integration-spec", handlers.HandleIntegrationRequest)

	return router
}
