package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterMetrics(router *gin.Engine) {
	handler := promhttp.Handler()
	fn := func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)
	}
	router.GET("/metrics", fn)
}
