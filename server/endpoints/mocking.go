package endpoints

import (
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func init() {
	if viper.GetBool("mock") {
		log.Info("Mocking enabled")
	}
}

func MockNotify() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Info("Mock endpoint handler", "method", ctx.Request.Method, "url", ctx.Request.URL.Path)
	}
}
