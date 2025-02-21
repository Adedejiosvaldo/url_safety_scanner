package main

import (
	"log"

	"github.com/adedejiosvaldo/safe_url/config"
	"github.com/adedejiosvaldo/safe_url/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// Struct to receive Telex messages
// type TelexMessage struct {
// 	Text string `json:"text"`
// }

// func modifyMessage(ctx *gin.Context) {
// 	var input TelexMessage
// 	if err := ctx.ShouldBindJSON(&input); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
// 		return
// 	}

// 	// Extract URLs from the message
// 	urls := cmd.ExtractURLs(input.Text)

// 	// Debug logging
// 	fmt.Printf("Received message: %s\n", input.Text)
// 	fmt.Printf("Extracted URLs: %v\n", urls)

// 	if len(urls) == 0 {
// 		fmt.Println("No URLs found in the message")
// 		ctx.JSON(http.StatusOK, gin.H{"text": input.Text}) // No URLs to scan
// 		return
// 	}

// 	// Scan each URL
// 	for _, url := range urls {
// 		fmt.Printf("Scanning URL: %s\n", url)
// 		isSafe, err := cmd.CheckURL(url)
// 		fmt.Printf("URL safety check result: %v, err: %v\n", isSafe, err)
// 		if err != nil {
// 			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "URL scanning failed", "details": err.Error()})
// 			return
// 		}
// 		if !isSafe {
// 			ctx.JSON(http.StatusOK, gin.H{"text": "⚠️ This message contains a potentially unsafe URL."})
// 			return
// 		}
// 	}

// 	// If all URLs are safe, return the message as-is
// 	ctx.JSON(http.StatusOK, gin.H{"text": input.Text})
// }

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"status": "ok"})
	})

	port := config.GetPort()
	router := routes.SetupRouter()

	log.Fatal(router.Run(":" + port))

}

// r.POST("/scan", func(ctx *gin.Context) {
// 	var input struct {
// 		URL string `json:"url" binding:"required"`
// 	}
// 	if err := ctx.ShouldBindJSON(&input); err != nil {
// 		ctx.JSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	result, err := cmd.UseURLScannerAPI(input.URL)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	isSafe, err := cmd.CheckURL(input.URL)
// 	if err != nil {
// 		ctx.JSON(500, gin.H{"error": err.Error()})
// 		return
// 	}

// 	ctx.JSON(200, gin.H{"message": "URL Safety Scanner", "result": result, "isSafe": isSafe})
// })
