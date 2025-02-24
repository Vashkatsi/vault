package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/Vashkatsi/vault/internal/application"
	"github.com/Vashkatsi/vault/internal/init"
)

var dataService *application.DataService

func init() {
	dataService = initdeps.InitializeDependencies()
}

func RegisterRoutes(router *gin.Engine) {
	router.GET("/health", healthCheck)
	router.GET("/ready", readinessCheck)
	router.POST("/store", storeData)
	router.POST("/retrieve", retrieveData)
	router.POST("/validate", validateData)
}

type StoreRequest struct {
	DataID string                 `json:"data_id"`
	Data   map[string]interface{} `json:"data"`
}

type RetrieveRequest struct {
	DataID string `json:"data_id"`
}

func storeData(context *gin.Context) {
	var req StoreRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	dataID, err := dataService.StoreData(req.DataID, req.Data)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data_id": dataID})
}

func retrieveData(context *gin.Context) {
	var req RetrieveRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	plainData, err := dataService.RetrieveData(req.DataID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": plainData})
}

func healthCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "healthy"})
}

func readinessCheck(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"status": "ready"})
}

func validateData(context *gin.Context) {
	context.JSON(http.StatusNotImplemented, gin.H{"error": "validate endpoint not implemented"})
}