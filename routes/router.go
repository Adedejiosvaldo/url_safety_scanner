package routes

import (
	"github.com/adedejiosvaldo/safe_url/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	// Add CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	router.POST("/scan-url", handlers.HandleScanRequest)
	router.GET("/ping", handlers.HealthCheck)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})
	router.GET("/integration-spec", handlers.HandleIntegrationRequest)

	return router
}
