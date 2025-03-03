package main

import (
	"github.com/Vashkatsi/vault/internal/config"
	"github.com/Vashkatsi/vault/internal/interface/api"
	"github.com/Vashkatsi/vault/internal/observability"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	observability.PrometheusMiddleware(router)

	cfg := config.LoadConfig()

	api.RegisterRoutes(router, cfg)

	router.Run(":" + cfg.Port)
}
