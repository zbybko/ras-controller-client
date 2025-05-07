package endpoints

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func DeviceInfoHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Debug("Getting device info")
		c.JSON(http.StatusOK, gin.H{
			"manufacturer":    viper.GetString("device.manufacturer"),
			"model":           viper.GetString("device.model"),
			"modelVersion":    viper.GetString("device.model-version"),
			"firmwareVersion": viper.GetString("device.firmware-version"),
		})
	}
}
