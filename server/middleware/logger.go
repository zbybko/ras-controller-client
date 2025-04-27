package middleware

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

var logger *log.Logger

var DefaultRequestLevel = log.InfoLevel

func Logger() gin.HandlerFunc {
	if logger == nil {
		logger = log.WithPrefix("GIN server")
		if viper.GetBool("debug") {
			logger.SetLevel(log.DebugLevel)
		}
	}

	return func(c *gin.Context) {
		c.Next()
		level := DefaultRequestLevel
		switch c.Writer.Status() {
		case http.StatusOK:
		default:
			level = log.ErrorLevel
		}
		logger.Log(level, "Request handled",
			"method", c.Request.Method,
			"url", c.Request.URL.Path,
			"status", c.Writer.Status())
	}
}
