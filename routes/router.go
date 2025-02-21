package routes

import (
	"github.com/adedejiosvaldo/safe_url/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/scan-url", handlers.HandleScanRequest)
	router.GET("/integration-spec", handlers.HandleIntegrationRequest)

	return router
}
