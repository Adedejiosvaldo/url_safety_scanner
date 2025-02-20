package main

import (
	"log"

	"github.com/adedejiosvaldo/safe_url/cmd"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	r.POST("/scan", func(ctx *gin.Context) {
		var input struct {
			URL string `json:"url" binding:"required"`
		}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}
		result, err := cmd.UseURLScannerAPI(input.URL)
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, gin.H{"message": "URL Safety Scanner", "result": result})
	})

	r.Run(":8080")

}
