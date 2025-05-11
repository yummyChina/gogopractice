package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func LogInfoMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log := logrus.New()
		log.SetFormatter(&logrus.JSONFormatter{})
		log.WithFields(logrus.Fields{
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Info("Request received")
		c.Next()
	}
}
