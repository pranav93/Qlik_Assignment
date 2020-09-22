package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pranav93/Qlik_Assignment/models"
)

// LatencyMiddleware LatencyMiddleware
func LatencyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		c.Next()

		latency := time.Since(t)
		log.Println(latency)
		log.Println(latency.Nanoseconds())
		// Running func in separate goroutine
		go models.LogLatency(latency.Nanoseconds())
	}
}
