package handlers

import (
	"net/http"

	"github.com/adedejiosvaldo/safe_url/models"
	"github.com/adedejiosvaldo/safe_url/services"
	"github.com/gin-gonic/gin"
)

// HandleScanRequest processes incoming messages and classifies URLs
func HandleScanRequest(c *gin.Context) {
	var msgReq models.Message

	if err := c.ShouldBindJSON(&msgReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Extract and classify URLs
	urls := services.ExtractURLs(msgReq.Message)
	urlClassifications := services.ClassifyURLs(urls)

	// Build response message
	response := models.ResponsePayload{
		EventName: "url_scanned",
		Message:   services.BuildResponseMessage(msgReq.Message, urlClassifications),
		URLs:      urls,
		Status:    "success",
		Username:  "url-scanner-bot",
	}

	c.JSON(http.StatusOK, response)
}

func HandleIntegrationRequest(c *gin.Context) {
	services.GetIntegrationSpec(c)

}
