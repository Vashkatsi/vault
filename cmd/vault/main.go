package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Vashkatsi/vault/internal/interface/api"
)

func main() {
	router := gin.Default()
	api.RegisterRoutes(router)
	router.Run(":8080")
}