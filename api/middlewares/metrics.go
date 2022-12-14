package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

var (
	apiRequestsCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_requests_total",
		Help: "API requests counter",
	}, []string{"method", "path", "status"})

	apiResponseTimeDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name: "api_response_time_duration",
		Help: "API response times",
		// adding 1ms, 2ms, 2.5ms + def buckets
		Buckets: append([]float64{.001, .002, .0025}, prometheus.DefBuckets...),
	}, []string{"method", "path", "status"})
)

func init() {
	prometheus.MustRegister(apiRequestsCounter)
	prometheus.MustRegister(apiResponseTimeDuration)
}

func Metrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()

		ctx.Next()

		var (
			elapsed  = time.Since(start)
			method   = ctx.Request.Method
			fullPath = ctx.FullPath()
			status   = fmt.Sprint(ctx.Writer.Status())
		)
		apiRequestsCounter.WithLabelValues(method, fullPath, status).Inc()
		apiResponseTimeDuration.WithLabelValues(method, fullPath, status).Observe(elapsed.Seconds())
	}
}
