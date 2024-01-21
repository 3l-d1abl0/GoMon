package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LoggerWithLogrus(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Log using Logrus

		start := time.Now()
		// Process the request
		c.Next()

		//latency := time.Now().Sub(start)
		latency := time.Since(start)

		//logger.WithFields
		logger.WithFields(logrus.Fields{
			"method":       c.Request.Method,
			"path":         c.Request.URL.Path,
			"latency":      latency,
			"status":       c.Writer.Status(),
			"ip":           c.Copy().ClientIP(),
			"query_params": c.Request.URL.Query(),
		}).Info("REQ")

		fmt.Println(latency, c.Writer.Status(), c.ClientIP())

	}
}
