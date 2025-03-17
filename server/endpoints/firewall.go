package endpoints

import (
	"net/http"
	"ras/management/firewall"

	"github.com/gin-gonic/gin"
)

func EnableFirewall(c *gin.Context) {
	if err := firewall.Enable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Firewall enabled"})
}

func DisableFirewall(c *gin.Context) {
	if err := firewall.Disable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Firewall disabled"})
}
