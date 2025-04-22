package server_management

import (
	"ras/config"

	"github.com/gin-gonic/gin"
)

func SetModeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		newMode := c.Param("mode")

		config.SetServerMode(newMode)
		c.JSON(200, gin.H{"mode": gin.Mode()})
	}
}
func GetModeHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"mode": config.GetServerMode()})
	}
}
