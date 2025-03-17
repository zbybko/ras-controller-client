package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ras/management/modems"
)

func EnableFirewall(c *gin.Context) {
	if err := modems.Enable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Firewall enabled"})
}

func DisableFirewall(c *gin.Context) {
	if err := modems.Disable(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "Firewall disabled"})
}
