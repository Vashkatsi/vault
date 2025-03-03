package observability

import (
	"github.com/gin-gonic/gin"
	ginprom "github.com/zsais/go-gin-prometheus"
)

func PrometheusMiddleware(router *gin.Engine) {
	prom := ginprom.NewPrometheus("gin")
	prom.Use(router)
}
