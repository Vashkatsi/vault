package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", healthCheck)
	router.GET("/ready", readinessCheck)
	router.POST("/store", storeData)
	router.POST("/retrieve", retrieveData)
	router.POST("/validate", validateData)
}

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func readinessCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func storeData(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "store endpoint not implemented"})
}

func retrieveData(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "retrieve endpoint not implemented"})
}

func validateData(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "validate endpoint not implemented"})
}